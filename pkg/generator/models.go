package imagen

import (
	"image/color"
	"math"
)

func Clamp(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func HSVToRGBA(h, s, v float64) color.RGBA {
	var i, f, p, q, t float64

	// Achromatic
	if s == 0 {
		outV := uint8(Clamp(v*255+0.5, 0, 255))
		return color.RGBA{outV, outV, outV, 0xFF}
	}

	h /= 60
	i = math.Floor(h)
	f = h - i
	p = v * (1 - s)
	q = v * (1 - s*f)
	t = v * (1 - s*(1-f))

	var r, g, b float64
	switch i {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		// r = q
		// g = v
		// b = p
		r = v
		g = t
		b = p
	case 2:
		// r = p
		// g = v
		// b = t
		r = v
		g = p
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	default:
		r = v // p looks cool
		g = p
		b = q
	}

	outR := uint8(Clamp(r*255+0.5, 0, 255)) 
	outG := uint8(Clamp(g*255+0.5, 0, 255))
	outB := uint8(Clamp(b*255+0.5, 0, 255)) 
	return color.RGBA{outR, outG, outB, 0xFF}
}
