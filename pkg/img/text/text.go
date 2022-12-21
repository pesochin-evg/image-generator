package text

import (
	"image"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
)

func AddRune(dst *image.RGBA, pattern image.Image, text string, point fixed.Point26_6) error {
	var fontSize float64 = 300
	data, err := os.ReadFile("pkg/img/text/Fonts/Arial Rounded Bold.ttf")
	if err != nil {
		log.Fatal("Font not found")
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

	_, err = c.DrawString(text, point)
	if err != nil {
		return err
	}

	return nil
}

func centerText(text string, fontSize float64, r image.Rectangle) fixed.Point26_6 {
	l := int(float64(len(text)) * fontSize * 0.55)
	h := int(fontSize * 0.6)
	return fixed.Point26_6{X: fixed.I((r.Dx() - l) / 2), Y: fixed.I(r.Dy() - (r.Dy()-h)/2)}
}
