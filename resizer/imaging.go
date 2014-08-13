package resizer

import (
	"image"

	"github.com/disintegration/imaging"
)

func imagingResize(src image.Image, w int, h int, method imaging.ResampleFilter) (image.Image, error) {
	return imaging.Resize(src, w, h, method), nil
}

func ImagingResizeNearestNeighbor(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.NearestNeighbor)
}

func ImagingResizeBox(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Box)
}

func ImagingResizeLinear(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Linear)
}

func ImagingResizeHermite(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Hermite)
}

func ImagingResizeMitchellNetravali(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.MitchellNetravali)
}

func ImagingResizeCatmullRom(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.CatmullRom)
}

func ImagingResizeBSpline(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.BSpline)
}

func ImagingResizeGaussian(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Gaussian)
}

func ImagingResizeBartlett(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Bartlett)
}

func ImagingResizeLanczos(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Lanczos)
}

func ImagingResizeHann(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Hann)
}

func ImagingResizeHamming(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Hamming)
}

func ImagingResizeWelch(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Welch)
}

func ImagingResizeCosine(src image.Image, w int, h int) (image.Image, error) {
	return imagingResize(src, w, h, imaging.Cosine)
}
