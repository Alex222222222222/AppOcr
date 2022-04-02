package screenshot

import (
	"image"
	"image/png"
	"os"
	"time"
)

func saveScreenShot(img *image.Image, saveDir string) (string, error) {

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
