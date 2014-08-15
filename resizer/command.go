package resizer

import (
	"bytes"
	"image"
	"io"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
)

type CommandResizeFunc func(c image.Config, inFormat string, outFormat string, w int, h int) (*exec.Cmd, error)

func (cr CommandResizeFunc) Name() string {
	name := runtime.FuncForPC(reflect.ValueOf(cr).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}

func (cr CommandResizeFunc) StreamResize(dst io.Writer, src io.Reader, format string, w int, h int) error {
	buf := bytes.NewBufferString("")
	cfg, inFormat, err := image.DecodeConfig(io.TeeReader(src, buf))
	if err != nil {
		return err
	}

	c, err := cr(cfg, inFormat, format, w, h)
	if err != nil {
		return err
	}

	c.Stdin = io.MultiReader(bytes.NewReader(buf.Bytes()), src)
	c.Stdout = dst
	return c.Run()
}

var CommandResizer = []CommandResizeFunc{
	ImageMagickResize,
	ImageMagickResizeWithDefine,
}
