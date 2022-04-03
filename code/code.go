package code

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/datamatrix"
	"github.com/makiuchi-d/gozxing/oned"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func ReadBarcode(img *image.Image) (string, error) {

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(*img)
	if err != nil {
		return "", err
	}

	reader := []gozxing.Reader{
		oned.NewCodaBarReader(),
		oned.NewCode128Reader(),
		oned.NewCode39Reader(),
		oned.NewCode93Reader(),
		oned.NewEAN13Reader(),
		oned.NewEAN8Reader(),
		oned.NewITFReader(),
		oned.NewUPCAReader(),
		oned.NewUPCEReader(),
	}

	errOut := err

	for i := 0; i < len(reader); i += 1 {
		oneReader := reader[i]
		res, err := oneReader.Decode(bmp, nil)
		if err != nil && !(strings.Contains(err.Error(), "NotFoundException")) {
			return "", err
		}
		if err == nil && len(res.GetText()) > 0 {
			return res.GetText(), nil
		}

		errOut = err

	}

	return "", errOut
}

func ReadQRCode(img *image.Image) (string, error) {

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(*img)
	if err != nil {
		return "", err
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}

func ReadDataMatrix(img *image.Image) (string, error) {

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(*img)
	if err != nil {
		return "", err
	}

	reader := datamatrix.NewDataMatrixReader()
	result, err := reader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}
