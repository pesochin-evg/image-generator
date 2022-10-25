package text

import (
	"image"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
)

func AddRune(dst *image.RGBA, pattern *image.RGBA, s string) error {
	point := fixed.Point26_6{X: fixed.I(30), Y: fixed.I(1200)}
	data, err := ioutil.ReadFile("/System/Library/Fonts/Supplemental/Times New Roman Bold.ttf")
	// data, err := ioutil.ReadFile("/System/Library/Fonts/Supplemental/Arial Rounded Bold.ttf")
	if err != nil {
		log.Fatal("File not found")
		return err
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		return err
	}

	var size float64 = 660

	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(pattern)
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetDPI(75)

	point2 := fixed.Point26_6{X: point.X - fixed.I(10), Y: point.Y}
	c2 := freetype.NewContext()
	c2.SetDst(dst)
	c2.SetClip(dst.Bounds())
	c2.SetSrc(image.White)
	c2.SetFont(f)
	c2.SetFontSize(size)
	c2.SetDPI(75)

	point3 := fixed.Point26_6{X: point.X + fixed.I(10), Y: point.Y }
	point4 := fixed.Point26_6{X: point.X, Y: point.Y - fixed.I(10)}
	point5 := fixed.Point26_6{X: point.X, Y: point.Y + fixed.I(10)}

	_, err = c2.DrawString(s, point2)
	if err != nil {
		return err
	}

	_, err = c2.DrawString(s, point3)
	if err != nil {
		return err
	}

	_, err = c2.DrawString(s, point4)
	if err != nil {
		return err
	}

	_, err = c2.DrawString(s, point5)
	if err != nil {
		return err
	}

	_, err = c.DrawString(s, point)
	if err != nil {
		return err
	}

	return nil
}

// func AddRune(dst *image.RGBA, pattern *image.RGBA, r rune) {
// 	point := fixed.Point26_6{X: fixed.I(100), Y: fixed.I(100)}
// 	f := opentype.NewFace(&sfnt.Font{}  goregular.TTF, nil)

// 	d := &font.Drawer{
// 		Dst:  dst,
// 		Src:  pattern,
// 		Face: f,
// 		Dot:  point,
// 	}
// 	d.DrawString(string(r))
// }
