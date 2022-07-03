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

func Generate(seed string) []*image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, Width, Height))
	var (
		IntSeed int64
		pow     int64
	)
	pow = 2

	for i := range seed {
		IntSeed += int64(i) * pow
		pow *= 2
	}

	f := field.New(IntSeed)

	step := Height / Blocks
	ch := make(chan bool)
	for i := 0; i < Blocks; i++ {
		go Fill(m, f, i*step, (i+1)*step, ch)
	}

	for i := 0; i < Blocks; i++ {
		<-ch
	}

	var result []*image.RGBA = make([]*image.RGBA, 3)
	result[0] = m
	result[1] = m.SubImage(image.Rectangle{image.Point{0, 0}, image.Point{828, 1720}}).(*image.RGBA)
	result[2] = m.SubImage(image.Rectangle{image.Point{828, 0}, image.Point{1656, 1720}}).(*image.RGBA)
	return result
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
				angle = 180 + tg
			}
		}
	}

	length := math.Sqrt(x*x + y*y)
	return HSVToRGBA(angle, length/max_length, 0.8)
}
