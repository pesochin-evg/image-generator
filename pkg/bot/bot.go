package bot

import (
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Antipascal/image-generator/pkg/field"
	imagen "github.com/Antipascal/image-generator/pkg/generator"

	tele "gopkg.in/telebot.v3"
)

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
		Ulist     = make(map[int64]User, 100)
		menu      = &tele.ReplyMarkup{ResizeKeyboard: true}
		empty     = &tele.ReplyMarkup{ResizeKeyboard: true}
		btnCreate = menu.Text("Create")
		admin     = &tele.User{
			ID:           541204191,
			FirstName:    "Evgeniy",
			LastName:     "Pesochin",
			Username:     "BRTL16",
			LanguageCode: "en",
			IsBot:        false,
		} //I know...
	)

	empty.RemoveKeyboard = true

	menu.Reply(
		menu.Row(btnCreate),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("This bot generates unique images especially for you and your significant other."+
			" Two wallpapers are created so that they form a common picture. The resulting drawing will be"+
			" related specifically to your couple, based on your names and the date you started the relationship.\n"+
			"To start just press create button", menu)
	})

	b.Handle(&btnCreate, func(c tele.Context) error {
		Ulist[c.Sender().ID] = User{State: 1}
		return c.Send("Enter your name:", empty)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		id := c.Sender().ID
		u, exists := Ulist[id]
		if !exists {
			return c.Send("To start just press create button", menu)
		}
		switch u.State {
		case 1:
			u.FName = c.Text()
			u.State = 2
			Ulist[id] = u
			return c.Send("Enter name of your significant other", empty)
		case 2:
			u.SName = c.Text()
			u.State = 3
			Ulist[id] = u
			return c.Send("Enter date when you started the relationship", empty)
		case 3:
			u.Date = c.Text()

			b.Send(admin, "* "+c.Sender().FirstName+" "+c.Sender().LastName+" - "+
				strconv.FormatInt(c.Sender().ID, 10)+" https://t.me/@"+c.Sender().Username+
				"( "+u.FName+" "+u.SName+" "+u.Date+" )")
			c.Send("Generating image for you (approx. 2 min)", empty)

			GenerateImages(id, u.FName+u.Date+u.SName)
			c.Send(&tele.Photo{File: tele.FromDisk(strconv.FormatInt(id, 10) + "_0.png")}, empty)
			f := &tele.Document{File: tele.FromDisk(strconv.FormatInt(id, 10) + "_1.png"), FileName: "First.png"}
			s := &tele.Document{File: tele.FromDisk(strconv.FormatInt(id, 10) + "_2.png"), FileName: "Second.png"}
			c.SendAlbum(tele.Album{f, s}, empty)
			DeleteImages(id)
			delete(Ulist, id)
			return c.Send("I hope you liked it", menu)
		}
		return c.Send("To start just press create button", menu)
	})

	b.Start()
}

func GenerateImages(id int64, seed string) {

	var (
		IntSeed int64
		pow     int64
	)
	pow = 2

	for i := range seed {
		IntSeed += int64(i) * pow
		pow *= 2
	}

	im := imagen.Generate(field.GenerateField(IntSeed), 1656, 1720)

	var result []*image.RGBA = make([]*image.RGBA, 3)
	result[0] = im
	result[1] = im.SubImage(image.Rectangle{image.Point{0, 0}, image.Point{828, 1720}}).(*image.RGBA)
	result[2] = im.SubImage(image.Rectangle{image.Point{828, 0}, image.Point{1656, 1720}}).(*image.RGBA)
	
	for i := 0; i < 3; i++ {
		f, err := os.Create(strconv.FormatInt(id, 10) + "_" + strconv.Itoa(i) + ".png")
		if err != nil {
			log.Println(err)
		}

		if png.Encode(f, result[i]) != nil {
			log.Println(err)
		}
		f.Close()
	}
}

func DeleteImages(id int64) {
	for i := 0; i < 3; i++ {
		e := os.Remove(strconv.FormatInt(id, 10) + "_" + strconv.Itoa(i) + ".png")
		if e != nil {
			log.Fatal(e)
		}
	}
}
