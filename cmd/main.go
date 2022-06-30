package main

import (
	"github.com/Antipascal/image-generator/pkg/generator"
	"log"
	"os"
	"image/png"
)

func main() {
	f, err := os.Create("outimage.png")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	m := imagen.Generate()
	if png.Encode(f, m) != nil {
		log.Println(err)
	}

}
