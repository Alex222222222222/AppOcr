package main

import (
	"encoding/base64"
	_ "image/jpeg"
	_ "image/png"

	"github.com/Alex222222222222/AppOcr/code"
	"github.com/Alex222222222222/AppOcr/config"
	"github.com/Alex222222222222/AppOcr/files"
	"github.com/Alex222222222222/AppOcr/notification"
	"github.com/Alex222222222222/AppOcr/ocr"
	"github.com/Alex222222222222/AppOcr/screenshot"
	"github.com/getlantern/systray"
	"golang.design/x/clipboard"
)

var c *config.UserConfig

func main() {
	var err error

	// read config if no user config exist, use default config
	c, err = config.ReadConfig("")
	if err != nil {
		panic(err)
	}

	// init the clipboard, see "golang.design/x/clipboard"
	err = clipboard.Init()
	if err != nil {
		panic(err)
	}

	// start the top bar item
	systray.Run(onReady, onExit)
}

func onReady() {
	// the app icon
	icon, err := base64.StdEncoding.DecodeString(files.Icon)
	if err != nil {
		panic(err)
	}
	systray.SetIcon(icon)

	// capture screenshot in interactive mode
	icon, err = base64.StdEncoding.DecodeString(files.ScreenshotIcon)
	if err != nil {
		panic(err)
	}
	bCaptureScreenShot := systray.AddMenuItem("Capture Screenshot", "Capture Screenshot")
	bCaptureScreenShot.SetIcon(icon)

	// scan qrcode
	icon, err = base64.StdEncoding.DecodeString(files.ScanQRCodeIcon)
	if err != nil {
		panic(err)
	}
	bQRCode := systray.AddMenuItem("Scan QRCode", "Scan QRCode")
	bQRCode.SetIcon(icon)

	// scan bar code
	icon, err = base64.StdEncoding.DecodeString(files.ScanBarCodeIcon)
	if err != nil {
		panic(err)
	}
	bBarCode := systray.AddMenuItem("Scan Bar Code", "Scan Bar Code")
	bBarCode.SetIcon(icon)

	// scan data matrix
	icon, err = base64.StdEncoding.DecodeString(files.ScanDataMatrixIcon)
	if err != nil {
		panic(err)
	}
	bDataMatrix := systray.AddMenuItem("Scan Data Matrix", "Scan Data Matrix")
	bDataMatrix.SetIcon(icon)

	// OCR
	icon, err = base64.StdEncoding.DecodeString(files.ScanOcrIcon)
	if err != nil {
		panic(err)
	}
	bOCR := systray.AddMenuItem("OCR", "OCR")
	bOCR.SetIcon(icon)

	// quit the app
	icon, err = base64.StdEncoding.DecodeString(files.QuitIcon)
	if err != nil {
		panic(err)
	}
	bQuit := systray.AddMenuItem("Quit", "Quit")
	bQuit.SetIcon(icon)

	go func() {
		for {
			select {
			case <-bQuit.ClickedCh:
				systray.Quit()
			case <-bCaptureScreenShot.ClickedCh:
				img, err := screenshot.CaptureScreenShot(c)

				if err != nil {
					err := notification.Notify(
						c,
						"Capture Screenshot failed",
						"",
						err.Error(),
					)
					if err != nil {
						panic(err)
					}
				} else {
					_, err := screenshot.SaveScreenShot(img, c.ScreenShotSaveDir)
					if err != nil {
						err := notification.Notify(
							c,
							"Capture Screenshot failed",
							"",
							err.Error(),
						)
						if err != nil {
							panic(err)
						}
					} else {
						err := notification.Notify(
							c,
							"Capture Screenshot successful",
							"",
							"Screenshot saved to "+c.ScreenShotSaveDir,
						)
						if err != nil {
							panic(err)
						}
					}
				}
			case <-bQRCode.ClickedCh:
				img, err := screenshot.CaptureScreenShot(c)

				if err != nil {
					err := notification.Notify(
						c,
						"Capture Screenshot failed",
						"",
						err.Error(),
					)
					if err != nil {
						panic(err)
					}
				} else {
					res, err := code.ReadQRCode(img) //readQRCode()
					if err != nil {
						err := notification.Notify(
							c,
							"Read QRCode failed",
							"",
							err.Error(),
						)
						if err != nil {
							panic(err)
						}
					} else {
						clipboard.Write(clipboard.FmtText, []byte(res))
						err := notification.Notify(
							c,
							"Read QRCode successful",
							"",
							"Result copied to clipboard",
						)
						if err != nil {
							panic(err)
						}
					}
				}
			case <-bBarCode.ClickedCh:
				img, err := screenshot.CaptureScreenShot(c)

				if err != nil {
					err := notification.Notify(
						c,
						"Capture Screenshot failed",
						"",
						err.Error(),
					)
					if err != nil {
						panic(err)
					}
				} else {
					res, err := code.ReadBarcode(img) //readQRCode()
					if err != nil {
						err := notification.Notify(
							c,
							"Read Bar Code failed",
							"",
							err.Error(),
						)
						if err != nil {
							panic(err)
						}
					} else {
						clipboard.Write(clipboard.FmtText, []byte(res))
						err := notification.Notify(
							c,
							"Read Bar Code successful",
							"",
							"Result copied to clipboard",
						)
						if err != nil {
							panic(err)
						}
					}
				}
			case <-bDataMatrix.ClickedCh:
				img, err := screenshot.CaptureScreenShot(c)

				if err != nil {
					err := notification.Notify(
						c,
						"Capture Screenshot failed",
						"",
						err.Error(),
					)
					if err != nil {
						panic(err)
					}
				} else {
					res, err := code.ReadDataMatrix(img) //readQRCode()
					if err != nil {
						err := notification.Notify(
							c,
							"Read Data Matrix failed",
							"",
							err.Error(),
						)
						if err != nil {
							panic(err)
						}
					} else {
						clipboard.Write(clipboard.FmtText, []byte(res))
						err := notification.Notify(
							c,
							"Read Data Matrix successful",
							"",
							"Result copied to clipboard",
						)
						if err != nil {
							panic(err)
						}
					}
				}
			case <-bOCR.ClickedCh:
				img, err := screenshot.CaptureScreenShot(c)

				if err != nil {
					err := notification.Notify(
						c,
						"Capture Screenshot failed",
						"",
						err.Error(),
					)
					if err != nil {
						panic(err)
					}
				} else {
					res, err := ocr.OcrMain(img, c) //readQRCode()
					if err != nil {
						err := notification.Notify(
							c,
							"OCR failed",
							"",
							err.Error(),
						)
						if err != nil {
							panic(err)
						}
					} else {
						clipboard.Write(clipboard.FmtText, []byte(res))
						err := notification.Notify(
							c,
							"OCR successful",
							"",
							"Result copied to clipboard",
						)
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}()
}

func onExit() {
	// clean up here
	// nothing to clean up actually
}
