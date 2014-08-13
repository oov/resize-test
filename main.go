package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/oov/resize-test/resizer"
)

func open(filename string) (image.Image, error) {
	f, err := os.Open("test.jpg")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	b := src.Bounds()
	r := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(r, r.Bounds(), src, b.Min, draw.Src)
	return r, nil
}

func savePNG(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func main() {
	img, err := open("test.jpg")
	if err != nil {
		panic(err)
	}

	w, h := img.Bounds().Dx()/7, img.Bounds().Dy()/7
	for idx, f := range resizer.Resizer {
		log.Println(f.Name(), "processing...")
		d, err := f(img, w, h)
		if err != nil {
			log.Println(f.Name(), err)
			continue
		}

		err = savePNG(fmt.Sprintf("output/%02d-%s.png", idx, f.Name()), d)
		if err != nil {
			log.Println(f.Name(), err)
			continue
		}
	}
}
