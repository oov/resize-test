package resizer

import (
	"errors"
	"image"
	"io"
	"reflect"
	"runtime"
	"strings"
)

type ResizeFunc func(src image.Image, w int, h int) (image.Image, error)

func (f ResizeFunc) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}

func (f ResizeFunc) StreamResize(dst io.Writer, src io.Reader, outFormat string, w int, h int) error {
	encoder, ok := encoders[outFormat]
	if !ok {
		return errors.New("unsupported image format: " + outFormat)
	}

	srcImg, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	dstImg, err := f(srcImg, w, h)
	if err != nil {
		return err
	}

	return encoder(dst, dstImg)
}

var Resizer = []ResizeFunc{
	NfntResizeNearestNeighbor,
	NfntResizeBilinear,
	NfntResizeBicubic,
	NfntResizeMitchellNetravali,
	NfntResizeLanczos2,
	NfntResizeLanczos3,
	GiftResizeNearestNeighbor,
	GiftResizeLinear,
	GiftResizeBox,
	GiftResizeCubic,
	GiftResizeLanczos,
	MoustachioResample,
	MoustachioResize,
	RezConvertBilinear,
	RezConvertBicubic,
	RezConvertLanczos2,
	RezConvertLanczos3,
	GraphicsScale,
	ImagingResizeNearestNeighbor,
	ImagingResizeBox,
	ImagingResizeLinear,
	ImagingResizeHermite,
	ImagingResizeMitchellNetravali,
	ImagingResizeCatmullRom,
	ImagingResizeBSpline,
	ImagingResizeGaussian,
	ImagingResizeBartlett,
	ImagingResizeLanczos,
	ImagingResizeHann,
	ImagingResizeHamming,
	ImagingResizeWelch,
	ImagingResizeCosine,
}
