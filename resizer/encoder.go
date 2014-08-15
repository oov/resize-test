package resizer

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

var encoders = map[string]func(dst io.Writer, src image.Image) error{
	"png":  png.Encode,
	"jpeg": func(dst io.Writer, src image.Image) error { return jpeg.Encode(dst, src, nil) },
	"gif":  func(dst io.Writer, src image.Image) error { return gif.Encode(dst, src, nil) },
}

func RegisterEncoderFormat(name string, encode func(dst io.Writer, src image.Image) error) {
	encoders[name] = encode
}
