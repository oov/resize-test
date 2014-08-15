package resizer

import (
	"fmt"
	"image"

	"github.com/bamiaux/rez"
)

func rezConvert(src image.Image, w int, h int, method rez.Filter) (image.Image, error) {
	r := image.Rect(0, 0, w, h)
	var m image.Image
	switch t := src.(type) {
	case *image.YCbCr:
		m = image.NewYCbCr(r, t.SubsampleRatio)
	case *image.RGBA:
		m = image.NewRGBA(r)
	case *image.NRGBA:
		m = image.NewNRGBA(r)
	case *image.Gray:
		m = image.NewGray(r)
	default:
		return nil, fmt.Errorf("unsupported image format")
	}
	err := rez.Convert(m, src, method)
	return m, err
}

func RezConvertBilinear(src image.Image, w int, h int) (image.Image, error) {
	return rezConvert(src, w, h, rez.NewBilinearFilter())
}

func RezConvertBicubic(src image.Image, w int, h int) (image.Image, error) {
	return rezConvert(src, w, h, rez.NewBicubicFilter())
}

func RezConvertLanczos2(src image.Image, w int, h int) (image.Image, error) {
	return rezConvert(src, w, h, rez.NewLanczosFilter(2.0))
}

func RezConvertLanczos3(src image.Image, w int, h int) (image.Image, error) {
	return rezConvert(src, w, h, rez.NewLanczosFilter(3.0))
}
