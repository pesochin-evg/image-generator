package imagen

import (
	"github.com/Antipascal/image-generator/pkg/field"
	"image"
	"image/color"
	"math"
)

const (
	Width  = 828
	Height = 1792
)

func Generate() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, Width, Height))
	f := field.New()
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {
			m.SetRGBA(x, y, AngleColor(f.Get(float64(x), float64(y), Height, Width)))
		}
	}
	return m
}

func AngleColor(v field.Vector, max_length float64) color.RGBA {
	var angle float64
	x, y := v.X, v.Y

	if y == 0 {
		if x > 0 {
			angle = 90
		} else {
			angle = 270
		}
	} else {
		var tg float64 = math.Atan(math.Abs(x/y)) * 180 / math.Pi
		if x >= 0 {
			if y > 0 {
				angle = tg
			} else {
				angle = 180 - tg
			}
		} else {
			if y > 0 {
				angle = 360 - tg
			} else {
				angle = 180 + tg
			}
		}
	}

	length := math.Sqrt(x*x + y*y)
	return HSVToRGBA(angle, length/max_length, 0.8)
}
