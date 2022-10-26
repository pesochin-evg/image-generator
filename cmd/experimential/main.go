package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/Antipascal/image-generator/pkg/img/field"
	filter "github.com/Antipascal/image-generator/pkg/img/filters"
	imagen "github.com/Antipascal/image-generator/pkg/img/generator"
	"github.com/Antipascal/image-generator/pkg/img/text"
)

func main() {
	// c := " "
	// i := int64(100)
	// for c != "q" {
		// fmt.Println(i)
		var seed int64 = 118
		// i++
		f1 := field.GenerateFieldFreq(seed, 200)
		f2 := field.GenerateFieldFreq(seed+118, 160)
		m := imagen.Generate(f1, 1000, 1000)
		// m = filter.BoxBlur(m, &m.Rect, 25)
		m = filter.Filter(m, &m.Rect, filter.BlueFilter)
		nm := imagen.Generate(f2, 400, 500)
		// nm = filter.BoxBlur(nm, &nm.Rect, 10)
		nm = filter.Filter(nm, &nm.Rect, filter.OrangeFilter)

		f, err := os.Create("output.jpg")
		if err != nil {
			fmt.Print(err)
		}
		defer f.Close()

		text.AddRune(m, nm, "test")

		if png.Encode(f, m) != nil {
			log.Println(err)
		}
		// fmt.Scan(&c)
	}
// }
//118 134
