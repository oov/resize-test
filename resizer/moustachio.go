package resizer

import (
	"image"

	"code.google.com/p/appengine-go/example/moustachio/resize"
)

func MoustachioResample(src image.Image, w int, h int) (image.Image, error) {
	return resize.Resample(src, src.Bounds(), w, h), nil
}

func MoustachioResize(src image.Image, w int, h int) (image.Image, error) {
	return resize.Resize(src, src.Bounds(), w, h), nil
}
