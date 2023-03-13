package icons

import (
	"encoding/base64"
	"errors"

	"github.com/Alex222222222222/AppOcr/files"
)

type IconType int64

const (
	IconTypeApp IconType = iota
	IconTypeCaptureScreenShot
	IconTypeScanQRCode
	IconTypeOCR
	IconTypeQuit
	IconTypeChangeTheme
	IconTypeScanBarCode
	IconTYpeScanDataMatrix
)

func GetIcon(iconType IconType, darkMode bool) ([]byte, error) {
	iconRaw, err := GetIconRaw(iconType, darkMode)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(iconRaw)
}

func GetIconRaw(iconType IconType, darkMode bool) (string, error) {
	switch iconType {
	case IconTypeApp:
		if darkMode {
			return files.IconDark, nil
		} else {
			return files.Icon, nil
		}
	case IconTypeCaptureScreenShot:
		if darkMode {
			return files.ScreenshotIconDark, nil
		} else {
			return files.ScreenshotIcon, nil
		}
	case IconTypeScanQRCode:
		if darkMode {
			return files.ScanQRCodeIconDark, nil
		} else {
			return files.ScanQRCodeIcon, nil
		}
	case IconTypeOCR:
		if darkMode {
			return files.ScanOcrIconDark, nil
		} else {
			return files.ScanOcrIcon, nil
		}
	case IconTypeQuit:
		if darkMode {
			return files.QuitIconDark, nil
		} else {
			return files.QuitIcon, nil
		}

	case IconTypeChangeTheme:
		if darkMode {
			return files.ChangeThemeIconDark, nil
		} else {
			return files.ChangeThemeIcon, nil
		}
	case IconTypeScanBarCode:
		if darkMode {
			return files.ScanBarCodeIconDark, nil
		} else {
			return files.ScanBarCodeIcon, nil
		}
	case IconTYpeScanDataMatrix:
		if darkMode {
			return files.ScanDataMatrixIconDark, nil
		} else {
			return files.ScanDataMatrixIcon, nil
		}
	default:
		return "", errors.New("invalid icon type")
	}
}
