package resizer

import (
	"image"

	"golang.org/x/image/draw"
)

func drawResize(src image.Image, w int, h int, q draw.Interpolator) (image.Image, error) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	q.Scale(m, m.Bounds(), src, src.Bounds(), draw.Over, nil)
	return m, nil
}

func DrawResizeNearestNeighbor(src image.Image, w int, h int) (image.Image, error) {
	return drawResize(src, w, h, draw.NearestNeighbor)
}

func DrawResizeApproxBiLinear(src image.Image, w int, h int) (image.Image, error) {
	return drawResize(src, w, h, draw.ApproxBiLinear)
}

func DrawResizeBiLinear(src image.Image, w int, h int) (image.Image, error) {
	return drawResize(src, w, h, draw.BiLinear)
}

func DrawResizeCatmullRom(src image.Image, w int, h int) (image.Image, error) {
	return drawResize(src, w, h, draw.CatmullRom)
}
