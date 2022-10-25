package filter

import "image/color"

func RedFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		uint8(float64(a)*1.0) + 30,
		uint8(float64(a) * 0.3),
		uint8(float64(a) * 0.4),
		255}
}

func GreyFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		uint8(float64(a)*1) - 20,
		uint8(float64(a) * 1),
		uint8(float64(a)*1) + 40,
		255}
}

func avg(c color.RGBA) uint8 {
	return uint8(0.299*float64(c.R) + 0.578*float64(c.G) + 0.114*float64(c.B))
}
