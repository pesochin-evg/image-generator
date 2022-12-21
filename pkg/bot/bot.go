package bot

import (
	"bytes"
	"github.com/Antipascal/image-generator/pkg/img/text"
	"golang.org/x/image/math/fixed"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	"github.com/Antipascal/image-generator/pkg/img/field"
	imagen "github.com/Antipascal/image-generator/pkg/img/generator"

	tele "gopkg.in/telebot.v3"
)

// Start bot with token IMAGEN_TOKEN from OS environment variables
func Start() {
	token, exists := os.LookupEnv("IMAGEN_TOKEN")
	if !exists {
		log.Println("No token found")
		return
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		Ulist     = make(map[int64]*User, 100)
		menu      = &tele.ReplyMarkup{ResizeKeyboard: true}
		empty     = &tele.ReplyMarkup{ResizeKeyboard: true}
		btnCreate = menu.Text("✨ Create ✨")
	)

	empty.RemoveKeyboard = true

	menu.Reply(
		menu.Row(btnCreate),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hey there! ☺️ We're Imagen, a bot that creates beautiful, "+
			"personalized images for couples. Whether you're celebrating your first "+
			"date or your 50th anniversary, we've got you covered.\n\nJust press the "+
			"'create' button and we'll generate two wallpapers that form a single, "+
			"unique picture just for you and your significant other. It's a fun and "+
			"easy way to celebrate your love and create something special together. 💜\n\n", menu)
	})

	b.Handle(&btnCreate, func(c tele.Context) error {
		Ulist[c.Sender().ID] = &User{State: FirstNameInput}
		return c.Send("1️⃣ To create your personalized image, we need to know your name. "+
			"\nPlease enter your name below:", empty)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		id := c.Sender().ID
		u, exists := Ulist[id]
		if !exists {
			return c.Send("To start just press create button 👇️", menu)
		}
		switch u.State {
		case FirstNameInput:
			return getFirstName(u, c, empty)
		case SecondNameInput:
			return getSecondName(u, c, empty)
		case DateInput:
			err := getDate(u, c)
			if err != nil {
				return err
			}
			defer delete(Ulist, id)
			return sendImages(u, c, menu)
		}
		return c.Send("To start just press create button 👇️", menu)
	})

	b.Start()
}

// GenerateImages Generate one big image and devices it into two parts
// images are stored at id_0.png id_1.png id_2.png
func GenerateImages(user User) ([]*image.RGBA, error) {

	seed := user.String()

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

	err := text.AddRune(
		im,
		image.White,
		string(user.FName[0]),
		fixed.Point26_6{X: fixed.I(90), Y: fixed.I(1630)})
	if err != nil {
		return nil, err
	}

	err = text.AddRune(
		im,
		image.White,
		string(user.SName[0]),
		fixed.Point26_6{X: fixed.I(1370), Y: fixed.I(1630)})
	if err != nil {
		return nil, err
	}

	var result = make([]*image.RGBA, 3)
	result[0] = im
	result[1] = im.SubImage(image.Rectangle{Min: image.Point{}, Max: image.Point{X: 828, Y: 1792}}).(*image.RGBA)
	result[2] = im.SubImage(image.Rectangle{Min: image.Point{X: 828}, Max: image.Point{X: 1656, Y: 1792}}).(*image.RGBA)

	return result, nil
}

func getFirstName(user *User, c tele.Context, opts ...interface{}) error {
	user.FName = c.Text()
	user.State = SecondNameInput
	return c.Send("2️⃣ Now, we need to know the name of your significant "+
		"other. \nPlease enter their name below.", opts...)
}

func getSecondName(user *User, c tele.Context, opts ...interface{}) error {
	user.SName = c.Text()
	user.State = DateInput
	return c.Send("3️⃣ Great! The final piece of information we need is the "+
		"date you started your relationship. \nPlease enter the date in the "+
		"format MM/DD/YYYY:", opts...)
}

func getDate(user *User, c tele.Context, opts ...interface{}) error {
	user.Date = c.Text()
	return nil
}

func sendImages(user *User, c tele.Context, opts ...interface{}) error {
	err := GetAdmins().Notify(GetUserString(c)+"( "+user.String()+" )", c)
	if err != nil {
		return err
	}

	err = c.Send("✅ Generating image for you")
	if err != nil {
		return err
	}

	result, err := GenerateImages(*user)
	if err != nil {
		return err
	}

	err = c.Send(&tele.Photo{File: tele.FromReader(readImage(result[0]))})
	if err != nil {
		return err
	}
	err = c.SendAlbum(tele.Album{
		&tele.Document{File: tele.FromReader(readImage(result[1])), FileName: "First.png"},
		&tele.Document{File: tele.FromReader(readImage(result[2])), FileName: "Second.png"},
	})
	if err != nil {
		return err
	}

	return c.Send("We hope you liked it \U0001FAF6", opts...)
}

func readImage(i *image.RGBA) io.Reader {
	buff := new(bytes.Buffer)
	err := png.Encode(buff, i)
	if err != nil {
		log.Println("failed to create buffer", err)
	}
	return bytes.NewReader(buff.Bytes())
}
