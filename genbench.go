// +build ignore

// This program generates benchmark_test.go
// Invoke as
//
//	go run genbench.go | gofmt >benchmark_test.go
package main

import (
	"log"
	"os"
	"text/template"

	"github.com/oov/resize-test/resizer"
)

func main() {
	t := template.Must(template.New("").Parse(program))
	err := t.Execute(os.Stdout, struct {
		Resizer        []resizer.ResizeFunc
		CommandResizer []resizer.CommandResizeFunc
	}{
		Resizer:        resizer.Resizer,
		CommandResizer: resizer.CommandResizer,
	})
	if err != nil {
		log.Fatal(err)
	}
}

var program = `// DO NOT EDIT.
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

{{range .Resizer}}
func BenchmarkScaler{{.Name}}(b *testing.B) {
	resizeFuncBench(b, resizer.{{.Name}})
}

{{end}}
{{range .Resizer}}
func BenchmarkFlow{{.Name}}(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.ResizeFunc(resizer.{{.Name}})))
}

{{end}}
{{range .CommandResizer}}
func BenchmarkFlow{{.Name}}(b *testing.B) {
	streamResizeFuncBench(b, resizer.StreamResizer(resizer.CommandResizeFunc(resizer.{{.Name}})))
}

{{end}}
`
