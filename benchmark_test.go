// DO NOT EDIT.
// Generate with: go run genbench.go | gofmt >benchmark_test.go

package main

import (
	"image"
	"testing"

	"github.com/oov/resize-test/resizer"
)

var img = image.NewRGBA(image.Rect(0, 0, 1000, 1000))

func resizeFuncBench(b *testing.B, f resizer.ResizeFunc) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f(img, 300, 300)
	}
	b.StopTimer()
}

func BenchmarkNfntResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeNearestNeighbor)
}

func BenchmarkNfntResizeBilinear(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeBilinear)
}

func BenchmarkNfntResizeBicubic(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeBicubic)
}

func BenchmarkNfntResizeMitchellNetravali(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeMitchellNetravali)
}

func BenchmarkNfntResizeLanczos2(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeLanczos2)
}

func BenchmarkNfntResizeLanczos3(b *testing.B) {
	resizeFuncBench(b, resizer.NfntResizeLanczos3)
}

func BenchmarkGiftResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeNearestNeighbor)
}

func BenchmarkGiftResizeLinear(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeLinear)
}

func BenchmarkGiftResizeBox(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeBox)
}

func BenchmarkGiftResizeCubic(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeCubic)
}

func BenchmarkGiftResizeLanczos(b *testing.B) {
	resizeFuncBench(b, resizer.GiftResizeLanczos)
}

func BenchmarkMoustachioResample(b *testing.B) {
	resizeFuncBench(b, resizer.MoustachioResample)
}

func BenchmarkMoustachioResize(b *testing.B) {
	resizeFuncBench(b, resizer.MoustachioResize)
}

func BenchmarkRezConvertBilinear(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertBilinear)
}

func BenchmarkRezConvertBicubic(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertBicubic)
}

func BenchmarkRezConvertLanczos2(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertLanczos2)
}

func BenchmarkRezConvertLanczos3(b *testing.B) {
	resizeFuncBench(b, resizer.RezConvertLanczos3)
}

func BenchmarkGraphicsScale(b *testing.B) {
	resizeFuncBench(b, resizer.GraphicsScale)
}

func BenchmarkImagingResizeNearestNeighbor(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeNearestNeighbor)
}

func BenchmarkImagingResizeBox(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBox)
}

func BenchmarkImagingResizeLinear(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeLinear)
}

func BenchmarkImagingResizeHermite(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHermite)
}

func BenchmarkImagingResizeMitchellNetravali(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeMitchellNetravali)
}

func BenchmarkImagingResizeCatmullRom(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeCatmullRom)
}

func BenchmarkImagingResizeBSpline(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBSpline)
}

func BenchmarkImagingResizeGaussian(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeGaussian)
}

func BenchmarkImagingResizeBartlett(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeBartlett)
}

func BenchmarkImagingResizeLanczos(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeLanczos)
}

func BenchmarkImagingResizeHann(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHann)
}

func BenchmarkImagingResizeHamming(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeHamming)
}

func BenchmarkImagingResizeWelch(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeWelch)
}

func BenchmarkImagingResizeCosine(b *testing.B) {
	resizeFuncBench(b, resizer.ImagingResizeCosine)
}
