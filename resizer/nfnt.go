package resizer

import (
	"image"

	"github.com/nfnt/resize"
)

func nfntResize(src image.Image, w int, h int, method resize.InterpolationFunction) (image.Image, error) {
	return resize.Resize(uint(w), uint(h), src, method), nil
}

func NfntResizeNearestNeighbor(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.NearestNeighbor)
}

func NfntResizeBilinear(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.Bilinear)
}

func NfntResizeBicubic(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.Bicubic)
}

func NfntResizeMitchellNetravali(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.MitchellNetravali)
}

func NfntResizeLanczos2(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.Lanczos2)
}

func NfntResizeLanczos3(src image.Image, w int, h int) (image.Image, error) {
	return nfntResize(src, w, h, resize.Lanczos3)
}
