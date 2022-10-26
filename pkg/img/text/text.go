package text

import (
	"image"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
)

func AddRune(dst *image.RGBA, pattern *image.RGBA, text string) error {
	var fontSize float64 = 190
	point := centerText(text, fontSize, dst.Bounds())
	// data, err := ioutil.ReadFile("/System/Library/Fonts/Supplemental/Times New Roman Bold.ttf")
	data, err := ioutil.ReadFile("/System/Library/Fonts/Supplemental/Arial Rounded Bold.ttf")
	if err != nil {
		log.Fatal("File not found")
		return err
	}
	f, err := freetype.ParseFont(data)
	if err != nil {
		return err
	}

	c := freetype.NewContext()
	c.SetDst(dst)
	c.SetClip(dst.Bounds())
	c.SetSrc(pattern)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetDPI(75)

	point2 := fixed.Point26_6{X: point.X - fixed.I(0), Y: point.Y}
	c2 := freetype.NewContext()
	c2.SetDst(dst)
	c2.SetClip(dst.Bounds())
	c2.SetSrc(image.White)
	c2.SetFont(f)
	c2.SetFontSize(fontSize)
	c2.SetDPI(75)

	_, err = c2.DrawString(text, point2)
	if err != nil {
		return err
	}


	_, err = c.DrawString(text, point)
	if err != nil {
		return err
	}

	return nil
}

func centerText(text string, fontSize float64, r image.Rectangle) fixed.Point26_6{
	l := int(float64(len(text)) * fontSize * 0.55)
	h := int(fontSize * 0.6)
	return fixed.Point26_6{X: fixed.I((r.Dx() - l) / 2), Y: fixed.I(r.Dy() - (r.Dy() - h) / 2)}
}
