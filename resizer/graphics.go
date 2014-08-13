package resizer

import (
	"image"

	"code.google.com/p/graphics-go/graphics"
)

func GraphicsScale(src image.Image, w int, h int) (image.Image, error) {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	err := graphics.Scale(m, src)
	return m, err
}
