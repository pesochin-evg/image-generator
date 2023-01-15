package bot

import (
	"bytes"
	"github.com/Antipascal/image-generator/pkg/FilterType"
	"github.com/Antipascal/image-generator/pkg/img/field"
	filter "github.com/Antipascal/image-generator/pkg/img/filters"
	imagen "github.com/Antipascal/image-generator/pkg/img/generator"
	tele "gopkg.in/telebot.v3"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
)

func sendImages(c *UserContext, t tele.Context) error {
	err := GetAdmins().Notify(GetUserString(t)+"( "+c.String()+" )", t)
	if err != nil {
		return err
	}

	err = t.Send("✅ Generating image for you")
	if err != nil {
		return err
	}

	result, err := GenerateImages(c)
	if err != nil {
		return err
	}

	err = t.Send(&tele.Photo{File: tele.FromReader(readImage(result[0])), Caption: "We hope you liked it \U0001FAF6"},
		BaseMenu)
	if err != nil {
		return err
	}
	return t.SendAlbum(tele.Album{
		&tele.Document{File: tele.FromReader(readImage(result[1])), FileName: "Left.png"},
		&tele.Document{File: tele.FromReader(readImage(result[2])), FileName: "Right.png"},
	})
}

func readImage(i *image.RGBA) io.Reader {
	buff := new(bytes.Buffer)
	err := png.Encode(buff, i)
	if err != nil {
		log.Println("failed to create buffer", err)
	}
	return bytes.NewReader(buff.Bytes())
}

// GenerateImages Generate one big image and devices it into two parts
// images are stored in memory
func GenerateImages(c *UserContext) ([]*image.RGBA, error) {

	seed := c.String()

	var (
		IntSeed int64
		pow     int64
	)
	pow = 2

	for i := range seed {
		IntSeed += int64(i) * pow
		pow *= 2
	}

	im := imagen.Generate(field.GenerateField(IntSeed), 1656, 1792)
	ApplyFilter(c, im)

	var result = make([]*image.RGBA, 3)
	result[0] = im
	result[1] = im.SubImage(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 828, Y: 1792}}).(*image.RGBA)
	result[2] = im.SubImage(image.Rectangle{Min: image.Point{X: 828}, Max: image.Point{X: 1656, Y: 1792}}).(*image.RGBA)

	return result, nil
}

func ApplyFilter(c *UserContext, m *image.RGBA) {
	var method func(color.RGBA) color.RGBA
	method = FilterType.GetFilterMethod(c.Filter)
	if method == nil {
		return
	}
	*m = *filter.Filter(m, &m.Rect, method)
}
