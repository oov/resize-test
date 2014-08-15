package resizer

import "io"

type StreamResizeFunc func(dst io.Writer, src io.Reader, outFormat string, w int, h int) error

type StreamResizer interface {
	StreamResize(dst io.Writer, src io.Reader, outFormat string, w int, h int) error
	Name() string
}

var AllResizer []StreamResizer

func init() {
	for _, f := range Resizer {
		AllResizer = append(AllResizer, f)
	}
	for _, f := range CommandResizer {
		AllResizer = append(AllResizer, f)
	}
}
