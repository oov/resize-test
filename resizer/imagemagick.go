package resizer

import (
	"fmt"
	"image"
	"os/exec"
)

func ImageMagickResize(c image.Config, inFormat string, outFormat string, w int, h int) (*exec.Cmd, error) {
	return exec.Command(
		"convert",
		"-resize",
		fmt.Sprintf("%dx%d", w, h),
		"-",
		fmt.Sprintf("%s:-", outFormat),
	), nil
}

func ImageMagickResizeWithDefine(c image.Config, inFormat string, outFormat string, w int, h int) (*exec.Cmd, error) {
	if inFormat != "jpeg" {
		return ImageMagickResize(c, inFormat, outFormat, w, h)
	}

	return exec.Command(
		"convert",
		"-define",
		fmt.Sprintf("jpeg:size=%dx%d", w, h),
		"-resize",
		fmt.Sprintf("%dx%d", w, h),
		"-",
		fmt.Sprintf("%s:-", outFormat),
	), nil
}
