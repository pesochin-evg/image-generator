package filter

import (
	"image/color"
	"math"
)

// RedFilter changes color c to color in red palette
func RedFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a + 30),
		G: setToBounds(a * 0.3),
		B: setToBounds(a * 0.4),
		A: 255}
}

// GreyFilter changes color c to color in grey palette
func GreyFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a - 20),
		G: setToBounds(a),
		B: setToBounds(a + 40),
		A: 255}
}

// PurpleFilter changes color c to color in purple palette
func PurpleFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a*0.4 + 30),
		G: setToBounds(a*0.4 + 30),
		B: setToBounds(a*1 + 40),
		A: 255}
}

// BlueFilter changes color c to color in blue palette
func BlueFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a - 40),
		G: setToBounds(a),
		B: setToBounds(a*1.2 + 60),
		A: 255}
}

// DarkBlueFilter changes color c to color in blue palette
func DarkBlueFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a * 0.2),
		G: setToBounds(a * 0.4),
		B: setToBounds(a * 1.5),
		A: 255}
}

// YellowFilter changes color c to color in yellow palette
func YellowFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a*1 + 40),
		G: setToBounds(a*1 + 40),
		B: setToBounds(a * 0.1),
		A: 255}
}

// OrangeFilter changes color c to color in orange palette
func OrangeFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a*1 + 50),
		G: setToBounds(a*0.55 + 50),
		B: setToBounds(a*0.3 - 10),
		A: 255}
}

// GreenFilter changes color c to color in green palette
func GreenFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a*0.8 + 20),
		G: setToBounds(a*1.2 + 20),
		B: setToBounds(a*0.6 + 20),
		A: 255}
}

// WhiteFilter changes color c to color in white palette
func WhiteFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a * 1.5),
		G: setToBounds(a * 1.5),
		B: setToBounds(a * 1.5),
		A: 255}
}

// BlackFilter changes color c to color in black palette
func BlackFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		R: setToBounds(a * 0.4),
		G: setToBounds(a * 0.4),
		B: setToBounds(a * 0.4),
		A: 255}
}

// Fits value a to bounds [0..255]
func setToBounds(a float64) uint8 {
	return uint8(math.Min(math.Max(50, a), math.MaxUint8))
}

// returns coefficient of the color
func avg(c color.RGBA) float64 {
	return 0.299*float64(c.R) + 0.578*float64(c.G) + 0.114*float64(c.B)
}
