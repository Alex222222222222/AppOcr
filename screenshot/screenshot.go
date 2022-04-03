package screenshot

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"time"

	"github.com/Alex222222222222/AppOcr/config"
	"github.com/kbinani/screenshot"
)

func SaveScreenShot(img *image.Image, saveDir string) (string, error) {

	f, err := os.Create(saveDir + "/ScreenShot " + time.Now().Format("Mon Jan 2 15:04:05 MST 2006") + ".png")
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = png.Encode(f, *img)
	if err != nil {
		return "", err
	}

	return "Save ScreenShot Successful", nil
}

func CaptureWholeScreen() ([]*image.Image, error) {
	n := screenshot.NumActiveDisplays()

	imgs := make([]*image.Image, n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		err = png.Encode(buf, img)

		imgP, _, err := image.Decode(bytes.NewReader(buf.Bytes()))
		if err != nil {
			return nil, err
		}

		imgs[i] = &imgP

		fmt.Println(bounds)
	}

	return imgs, nil

}

func CaptureScreenShot(c *config.UserConfig) (*image.Image, error) {
	f1, err := os.Create(c.CacheDir + "/log.txt")
	if err != nil {
		return nil, err
	}
	defer f1.Close()

	cmd := exec.Command("screencapture", "-i", "-D", "1", c.CacheDir+"/screenshot.png")
	cmd.Stderr = f1
	cmd.Stdout = f1

	err = cmd.Run()
	if err != nil {
		errorMessage, errNew := os.ReadFile(c.CacheDir + "/log.txt")
		if errNew != nil {
			return nil, errNew
		}
		return nil, errors.New(string(errorMessage) + "\n" + err.Error())
	}

	f2, err := os.ReadFile(c.CacheDir + "/screenshot.png")
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bytes.NewReader(f2))
	return &img, err
}
