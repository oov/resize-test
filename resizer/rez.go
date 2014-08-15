package resizer

import (
	"image"

	"github.com/bamiaux/rez"
)

func rezConvert(src image.Image, w int, h int, method rez.Filter) (image.Image, error) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
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
