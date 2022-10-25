package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/Antipascal/image-generator/pkg/field"
	filter "github.com/Antipascal/image-generator/pkg/filters"
	imagen "github.com/Antipascal/image-generator/pkg/generator"
	"github.com/Antipascal/image-generator/pkg/text"
)

func main() {
	var seed int64 = 923469394
	f1 := field.GenerateField(seed)
	f2 := field.GenerateFieldFreq(seed/10, 250)
	m := imagen.Generate(f1, 3024, 1964)
	m = filter.BoxBlur(m, &m.Rect, 40)
	// m = filter.Filter(m, &m.Rect, filter.GreyFilter)
	nm := imagen.Generate(f2, 3000, 1000)
	// nm = filter.BoxBlur(nm, &nm.Rect, 10)
	// nm = filter.Filter(nm, &nm.Rect, filter.GreyFilter)

	f, err := os.Create("output.jpg")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	text.AddRune(m, nm, "Fly Dubai")

	if png.Encode(f, m) != nil {
		log.Println(err)
	}

}
