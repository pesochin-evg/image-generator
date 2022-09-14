package imagen

import (
	"image"
	"image/color"
	"math"

	"github.com/Antipascal/image-generator/pkg/field"
)

const (
	Width  = 1656
	Height = 1720
	Blocks = 10
)

func Fill(m *image.RGBA, f *field.Field, minY, maxY int, ch chan bool) {
	for x := 0; x < Width; x++ {
		for y := minY; y <= maxY; y++ {
			m.SetRGBA(x, y, AngleColor(f.Get(float64(x), float64(y), Height, Width)))
		}
	}
	ch <- true
}

func Generate(seed int64) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, Width, Height))

	f := field.New(seed)

	step := Height / Blocks
	ch := make(chan bool)
	for i := 0; i < Blocks; i++ {
		go Fill(m, f, i*step, (i+1)*step, ch)
	}

	for i := 0; i < Blocks; i++ {
		<-ch
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
				angle = 180 + tg
			}
		} else {
			if y > 0 {
				angle = 360 - tg
			} else {
				angle = 180 - tg
			}
		}
	}

	length := math.Sqrt(x*x + y*y)
	return HSVToRGBA(angle, length/max_length, 0.8)
}
