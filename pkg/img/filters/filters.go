package filter

import (
	"image/color"
	"math"
)

// changes color c to color in red colorpalette
func RedFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a + 30),
		setToBounds(a * 0.3),
		setToBounds(a * 0.4),
		255}
}

// changes color c to color in grey colorpalette
func GreyFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a - 20),
		setToBounds(a),
		setToBounds(a + 40),
		255}
}

// changes color c to color in purple colorpalette
func PurpleFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 0.4 + 30),
		setToBounds(a * 0.4 + 30),
		setToBounds(a * 1 + 40),
		255}
}

// changes color c to color in blue colorpalette
func BlueFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	// return color.RGBA{
	// 	setToBounds(a * 0.2 + 40),
	// 	setToBounds(a * 0.7 + 30),
	// 	setToBounds(a * 1.5 + 40),
	// 	255}
	return color.RGBA{
		setToBounds(a - 40),
		setToBounds(a),
		setToBounds(a * 1.2 + 60),
		255}
}

// changes color c to color in blue colorpalette
func DarkBlueFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 0.2),
		setToBounds(a * 0.4),
		setToBounds(a * 1.5),
		255}
}

// changes color c to color in yellow colorpalette
func YellowFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 1 + 40),
		setToBounds(a * 1 + 40),
		setToBounds(a * 0.1 ),
		255}
}

// changes color c to color in orange colorpalette
func OrangeFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 1 + 50),
		setToBounds(a * 0.55 + 50),
		setToBounds(a * 0.3 - 10),
		255}
}

// changes color c to color in green colorpalette
func GreenFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 0.8 + 20),
		setToBounds(a * 1.2 + 20),
		setToBounds(a * 0.6 + 20),
		255}
}

// changes color c to color in white colorpalette
func WhiteFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 1.5),
		setToBounds(a * 1.5),
		setToBounds(a * 1.5),
		255}
}

// changes color c to color in black colorpalette
func BlackFilter(c color.RGBA) color.RGBA {
	a := avg(c)
	return color.RGBA{
		setToBounds(a * 0.4),
		setToBounds(a * 0.4),
		setToBounds(a * 0.4),
		255}
}

// Fits value a to bounds [0..255]
func setToBounds(a float64) uint8 {
	return uint8(math.Min(math.Max(50, a), math.MaxUint8))
}

// returns coefficient of the color
func avg(c color.RGBA) float64 {
	return 0.299*float64(c.R) + 0.578*float64(c.G) + 0.114*float64(c.B)
}
