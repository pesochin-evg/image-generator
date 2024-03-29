package imagen

import (
	"image"
	"image/color"
	"math"

	"github.com/Antipascal/image-generator/pkg/img/field"
)

func Fill(m *image.RGBA, f *field.Field, minY, maxY int, ch chan bool) {
	width := m.Bounds().Dx()
	for x := 0; x < width; x++ {
		for y := minY; y <= maxY; y++ {
			m.SetRGBA(x, y, AngleColor(f.Get(x, y)))
		}
	}
	ch <- true
}

func Generate(f *field.Field, width, height int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	blocks := 1
	for i := 15; i > 0; i-- {
		if height%i == 0 {
			blocks = i
			break
		}
	}

	step := height / blocks
	ch := make(chan bool)
	for i := 0; i < blocks; i++ {
		go Fill(m, f, i*step, (i+1)*step, ch)
	}

	for i := 0; i < blocks; i++ {
		<-ch
	}

	return m
}

func AngleColor(v field.Vector, maxLength float64) color.RGBA {
	var angle float64
	x, y := v.X, v.Y

	if y == 0 {
		if x > 0 {
			angle = 90
		} else {
			angle = 270
		}
	} else {
		var tg = math.Atan(math.Abs(x/y)) * 180 / math.Pi
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
	return HSVToRGBA(angle, length/maxLength, 0.8)
}
