package resizer

import (
	"image"
	"reflect"
	"runtime"
	"strings"
)

type ResizeFunc func(src image.Image, w int, h int) (image.Image, error)

func (f ResizeFunc) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
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
