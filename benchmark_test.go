// DO NOT EDIT.
// Generate with: go run genbench.go | gofmt >benchmark_test.go

package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"testing"

	"github.com/oov/resize-test/resizer"
)

var imgBuf, img = loadImage()

func loadImage() ([]byte, image.Image) {
	buf, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		panic(err)
	}

	img, _, err = image.Decode(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}

	return buf, img
}

func resizeFuncBench(b *testing.B, f resizer.ResizeFunc) {
	for i := 0; i < b.N; i++ {
		f(img, 300, 300)
	}
}

func streamResizeFuncBench(b *testing.B, f resizer.StreamResizer) {
	var err error
	for i := 0; i < b.N; i++ {
		err = f.StreamResize(ioutil.Discard, bytes.NewReader(imgBuf), "png", 300, 300)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkScalerNfntResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeNearestNeighbor)
}

func BenchmarkScalerNfntResizeBilinear(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeBilinear)
}

func BenchmarkScalerNfntResizeBicubic(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeBicubic)
}

func BenchmarkScalerNfntResizeMitchellNetravali(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeMitchellNetravali)
}

func BenchmarkScalerNfntResizeLanczos2(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeLanczos2)
}

func BenchmarkScalerNfntResizeLanczos3(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeLanczos3)
}

func BenchmarkScalerGiftResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeNearestNeighbor)
}

func BenchmarkScalerGiftResizeLinear(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeLinear)
}

func BenchmarkScalerGiftResizeBox(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeBox)
}

func BenchmarkScalerGiftResizeCubic(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeCubic)
}

func BenchmarkScalerGiftResizeLanczos(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeLanczos)
}

func BenchmarkScalerMoustachioResample(b *testing.B) {
	resizeFuncBench(b, resizer.MoustachioResample)
}

func BenchmarkScalerMoustachioResize(b *testing.B) {
	resizeFuncBench(b, resizer.MoustachioResize)
}

func BenchmarkScalerRezConvertBilinear(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertBilinear)
}

func BenchmarkScalerRezConvertBicubic(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertBicubic)
}

func BenchmarkScalerRezConvertLanczos2(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertLanczos2)
}

func BenchmarkScalerRezConvertLanczos3(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertLanczos3)
}

func BenchmarkScalerGraphicsScale(b *testing.B) {
	resizeFuncBench(b, resizer.GraphicsScale)
}

func BenchmarkScalerImagingResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeNearestNeighbor)
}

func BenchmarkScalerImagingResizeBox(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBox)
}

func BenchmarkScalerImagingResizeLinear(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeLinear)
}

func BenchmarkScalerImagingResizeHermite(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHermite)
}

func BenchmarkScalerImagingResizeMitchellNetravali(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeMitchellNetravali)
}

func BenchmarkScalerImagingResizeCatmullRom(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeCatmullRom)
}

func BenchmarkScalerImagingResizeBSpline(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBSpline)
}

func BenchmarkScalerImagingResizeGaussian(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeGaussian)
}

func BenchmarkScalerImagingResizeBartlett(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBartlett)
}

func BenchmarkScalerImagingResizeLanczos(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeLanczos)
}

func BenchmarkScalerImagingResizeHann(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHann)
}

func BenchmarkScalerImagingResizeHamming(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHamming)
}

func BenchmarkScalerImagingResizeWelch(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeWelch)
}

func BenchmarkScalerImagingResizeCosine(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeCosine)
}

func BenchmarkFlowNfntResizeNearestNeighbor(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeNearestNeighbor)))
}

func BenchmarkFlowNfntResizeBilinear(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeBilinear)))
}

func BenchmarkFlowNfntResizeBicubic(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeBicubic)))
}

func BenchmarkFlowNfntResizeMitchellNetravali(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeMitchellNetravali)))
}

func BenchmarkFlowNfntResizeLanczos2(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeLanczos2)))
}

func BenchmarkFlowNfntResizeLanczos3(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.NfntResizeLanczos3)))
}

func BenchmarkFlowGiftResizeNearestNeighbor(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GiftResizeNearestNeighbor)))
}

func BenchmarkFlowGiftResizeLinear(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GiftResizeLinear)))
}

func BenchmarkFlowGiftResizeBox(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GiftResizeBox)))
}

func BenchmarkFlowGiftResizeCubic(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GiftResizeCubic)))
}

func BenchmarkFlowGiftResizeLanczos(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GiftResizeLanczos)))
}

func BenchmarkFlowMoustachioResample(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.MoustachioResample)))
}

func BenchmarkFlowMoustachioResize(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.MoustachioResize)))
}

func BenchmarkFlowRezConvertBilinear(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.RezConvertBilinear)))
}

func BenchmarkFlowRezConvertBicubic(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.RezConvertBicubic)))
}

func BenchmarkFlowRezConvertLanczos2(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.RezConvertLanczos2)))
}

func BenchmarkFlowRezConvertLanczos3(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.RezConvertLanczos3)))
}

func BenchmarkFlowGraphicsScale(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.GraphicsScale)))
}

func BenchmarkFlowImagingResizeNearestNeighbor(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeNearestNeighbor)))
}

func BenchmarkFlowImagingResizeBox(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeBox)))
}

func BenchmarkFlowImagingResizeLinear(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeLinear)))
}

func BenchmarkFlowImagingResizeHermite(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeHermite)))
}

func BenchmarkFlowImagingResizeMitchellNetravali(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeMitchellNetravali)))
}

func BenchmarkFlowImagingResizeCatmullRom(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeCatmullRom)))
}

func BenchmarkFlowImagingResizeBSpline(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeBSpline)))
}

func BenchmarkFlowImagingResizeGaussian(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeGaussian)))
}

func BenchmarkFlowImagingResizeBartlett(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeBartlett)))
}

func BenchmarkFlowImagingResizeLanczos(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeLanczos)))
}

func BenchmarkFlowImagingResizeHann(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeHann)))
}

func BenchmarkFlowImagingResizeHamming(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeHamming)))
}

func BenchmarkFlowImagingResizeWelch(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeWelch)))
}

func BenchmarkFlowImagingResizeCosine(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.ImagingResizeCosine)))
}

func BenchmarkFlowImageMagickResize(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.CommandResizeFunc(resizer.ImageMagickResize)))
}

func BenchmarkFlowImageMagickResizeWithDefine(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.CommandResizeFunc(resizer.ImageMagickResizeWithDefine)))
}
