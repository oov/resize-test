package resizer

import (
	"image"

	"github.com/disintegration/gift"
)

func giftResize(src image.Image, w int, h int, method gift.Resampling) (image.Image, error) {
	g := gift.New(gift.Resize(w, h, method))
	m := image.NewRGBA(g.Bounds(src.Bounds()))
	g.Draw(m, src)
	return m, nil
}

func GiftResizeNearestNeighbor(src image.Image, w int, h int) (image.Image, error) {
	return giftResize(src, w, h, gift.NearestNeighborResampling)
}

func GiftResizeLinear(src image.Image, w int, h int) (image.Image, error) {
	return giftResize(src, w, h, gift.LinearResampling)
}

func GiftResizeBox(src image.Image, w int, h int) (image.Image, error) {
	return giftResize(src, w, h, gift.BoxResampling)
}

func GiftResizeCubic(src image.Image, w int, h int) (image.Image, error) {
	return giftResize(src, w, h, gift.CubicResampling)
}

func GiftResizeLanczos(src image.Image, w int, h int) (image.Image, error) {
	return giftResize(src, w, h, gift.LanczosResampling)
}
