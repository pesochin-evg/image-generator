package filter

import (
	"image"
	"image/color"
)

// Filter add f filter function to every pixel on
// image m in r bounds.
func Filter(m *image.RGBA, r *image.Rectangle, f func(color.RGBA) color.RGBA) *image.RGBA {
	width := r.Dx()
	height := r.Dy()
	nm := image.NewRGBA(m.Rect)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			nm.SetRGBA(x, y, f(m.RGBAAt(x, y)))
		}
	}
	return nm
}

// BoxBlur add box blur on image m in r bounds with
// kernel size k.
func BoxBlur(m *image.RGBA, r *image.Rectangle, k int) *image.RGBA {
	nm := image.NewRGBA(m.Rect)
	for x := 0; x < m.Rect.Dx(); x++ {
		for y := 0; y < m.Rect.Dy(); y++ {
			if x >= r.Min.X && x < r.Max.X && y >= r.Min.Y && y < r.Max.Y {
				nm.SetRGBA(x, y, avgRectangle(m, r, x, y, k/2))
			} else {
				nm.SetRGBA(x, y, m.RGBAAt(x, y))
			}
		}
	}
	return nm
}

func avgRectangle(m *image.RGBA, re *image.Rectangle, x, y, size int) color.RGBA {
	var r, g, b, count uint32 = 0, 0, 0, 0
	for w := max(x-size, re.Min.X); w < min(x+size+1, re.Max.X); w++ {
		for h := max(y-size, re.Min.Y); h < min(y+size+1, re.Max.Y); h++ {
			c := m.RGBAAt(w, h)
			r, g, b = r+uint32(c.R), g+uint32(c.G), b+uint32(c.B)
			count++
		}
	}
	return color.RGBA{
		R: uint8(r / count),
		G: uint8(g / count),
		B: uint8(b / count),
		A: 255}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
