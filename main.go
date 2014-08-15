package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/oov/resize-test/resizer"
)

func save(f resizer.StreamResizer, r io.Reader, filename string, format string, w int, h int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return f.StreamResize(file, r, format, w, h)
}

func main() {
	imgBuf, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewReader(imgBuf)
	cfg, _, err := image.DecodeConfig(r)
	if err != nil {
		log.Fatalln(err)
	}

	md, err := os.Create("output/README.md")
	if err != nil {
		log.Fatalln(err)
	}
	defer md.Close()

	fmt.Fprintln(md, `# Image Quality Preview`)
	fmt.Fprintln(md)

	w, h := cfg.Width/7, cfg.Height/7
	for idx, f := range resizer.AllResizer {
		log.Println(f.Name(), "processing...")

		r.Seek(0, os.SEEK_SET)
		err := save(f, r, fmt.Sprintf("output/%02d-%s.png", idx, f.Name()), "png", w, h)
		if err != nil {
			log.Println(f.Name(), err)
			continue
		}

		fmt.Fprintf(md, "### %s\n\n", f.Name())
		fmt.Fprintf(md, "![%s output image](%02d-%s.png)\n\n", f.Name(), idx, f.Name())
	}
}
