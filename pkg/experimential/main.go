package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	imagen "github.com/Antipascal/image-generator/pkg/generator"
	filter "github.com/Antipascal/image-generator/pkg/filters"
)

func main() {
	var seed int64 = 999999
	m := imagen.Generate(seed, 3000, 2000)
	f, err := os.Create("output.jpg")
	if err != nil {
		fmt.Print(err)
	}

	defer f.Close()

	filter.BWFilter(m)

	if png.Encode(f, m) != nil {
		log.Println(err)
	}

}

