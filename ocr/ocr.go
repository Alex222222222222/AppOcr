package ocr

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"strings"

	"github.com/Alex222222222222/AppOcr/config"
	spacewrapper "github.com/ranghetto/go_ocr_space"
)

// I also cannot find pure golang ocr engine.
// There is a opensource Tesseract golang wrapper, but there were error, when I tried to compile the app
// Online space ocr work well, if you donot want to install tesseract. Register a space ocr api at the official website

func ocrWithSpaceOcr(img *image.Image, c *config.UserConfig) (string, error) {
	config := spacewrapper.InitConfig(c.SpaceOCRAPI, c.OcrLanguage[0].SpaceOCR)

	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, *img, nil)
	if err != nil {
		return "", err
	}

	result, err := config.ParseFromBase64("data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()))
	if err != nil {
		return "", err
	}

	return result.JustText(), nil

}

func ocrWithTesseract(img *image.Image, c *config.UserConfig) (string, error) {
	if _, err := os.Stat(c.CacheDir); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(c.CacheDir, os.ModePerm)
	}

	f, err := os.Create(c.CacheDir + "/tesseract.png")
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, *img)
	if err != nil {
		return "", err
	}

	f1, err := os.Create(c.CacheDir + "/log.txt")
	if err != nil {
		return "", err
	}

	languages := make([]string, len(c.OcrLanguage))

	for i := 0; i < len(c.OcrLanguage); i += 1 {
		languages[i] = c.OcrLanguage[i].Tesseract
	}

	cmd := exec.Command(c.TesseractExecutablePath, c.CacheDir+"/tesseract.png", c.CacheDir+"/res", "-l", strings.Join(languages, "+"))
	cmd.Stderr = f1
	cmd.Stdout = f1
	err = cmd.Run()
	if err != nil {
		res, errNew := os.ReadFile(c.CacheDir + "/log.txt")
		if errNew != nil {
			return string(res), errNew
		}
		return string(res), err
	}

	res, err := os.ReadFile(c.CacheDir + "/res.txt")
	return string(res), err

}

func OcrMain(img *image.Image, c *config.UserConfig) (string, error) {
	switch c.OCREngine {
	case "SpaceOCR":
		return ocrWithSpaceOcr(img, c)
	case "Tesseract":
		return ocrWithTesseract(img, c)
	}

	return "", errors.New("Specified ocr engine does not supported.")
}
