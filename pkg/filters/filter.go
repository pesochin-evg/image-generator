package filters

import (
	"image"
	"image/color"
)

func BWFilter(m *image.RGBA) {
	width := m.Bounds().Dx()
	height := m.Bounds().Dy()
	for x := 0; x < width/2; x++ {
		for y := 0; y <= 2*height/4; y++ {
			a := avg(m.RGBAAt(x, y))
			m.SetRGBA(x, y, color.RGBA{
				uint8(float64(a)*1.0) + 30,
				uint8(float64(a) * 0.3),
				uint8(float64(a) * 0.4),
				255})
		}
	}
	// for y := height / 4; y <= 2*height/4; y++ {
	// 	a := avg(m.RGBAAt(x, y))
	// 	m.SetRGBA(x, y, color.RGBA{
	// 		uint8(float64(a)*1.0) + 10,
	// 		uint8(float64(a) * 1.1),
	// 		uint8(float64(a) * 0.2),
	// 		255})
	// }
	for x := width / 2; x < width; x++ {
		for y := 0; y <= height; y++ {
			a := avg(m.RGBAAt(x, y))
			m.SetRGBA(x, y, color.RGBA{
				uint8(float64(a) * 1) - 20,
				uint8(float64(a) * 1),
				uint8(float64(a) * 1.0) + 40,
				255})
		}
	}
}

func avg(c color.RGBA) uint8 {
	return uint8(0.299*float64(c.R) + 0.578*float64(c.G) + 0.114*float64(c.B))
}

// m.SetRGBA(x, y, color.RGBA{
// 	uint8(float64(a) * 1.2),
// 	uint8(float64(a) * 0.3),
// 	uint8(float64(a) * 0.3),
// 	255})
