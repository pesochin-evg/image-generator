package imagen

import (
	// "fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

const (
	Width  = 1024
	Height = 1024
)

func Generate() {
	m := image.NewRGBA(image.Rect(0, 0, Width, Height))
	var newX, newY float64
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			// if math.Pow(float64(x-Width/2), 2)+math.Pow(float64(y-Height/2), 2) >= 400 {
			newX = float64(x - Width/2)
			newY = float64(Height/2 - y)
			if newY == 0 {
				newY = 0.00001
			}
			if newY >= 0 {
				m.SetRGBA(x, y, HSVToRGBA(math.Atan(newX/newY)*180/math.Pi+45, 1, 1))
			} else {
				m.SetRGBA(x, y, HSVToRGBA(math.Atan(newX/newY)*180/math.Pi+225, 1, 1))
			}
			// fmt.Println(math.Atan(newX/newY)*180/math.Pi + 180)
			// }
		}
	}

	f, err := os.Create("outimage.png")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		log.Println(err)
	}

}
