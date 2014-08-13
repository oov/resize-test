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
	err := t.Execute(os.Stdout, resizer.Resizer)
	if err != nil {
		log.Fatal(err)
	}
}

var program = `// DO NOT EDIT.
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

{{range .}}
func Benchmark{{.Name}}(b *testing.B) {
	resizeFuncBench(b, resizer.{{.Name}})
}

{{end}}
`
