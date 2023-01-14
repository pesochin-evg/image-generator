package bot

import (
	"github.com/Antipascal/image-generator/pkg/FilterType"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	BaseMenu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				tele.ReplyButton{Text: "✨ Create ✨"},
			},
		},
	}

	FilterMenu = &tele.ReplyMarkup{
		ResizeKeyboard: true,
		ReplyKeyboard: [][]tele.ReplyButton{
			{
				tele.ReplyButton{Text: FilterType.None},
				tele.ReplyButton{Text: FilterType.Red},
				tele.ReplyButton{Text: FilterType.Grey},
				tele.ReplyButton{Text: FilterType.Purple},
				tele.ReplyButton{Text: FilterType.DarkBlue},
			},
			{
				tele.ReplyButton{Text: FilterType.Yellow},
				tele.ReplyButton{Text: FilterType.Orange},
				tele.ReplyButton{Text: FilterType.Green},
				tele.ReplyButton{Text: FilterType.White},
				tele.ReplyButton{Text: FilterType.Black},
			},
		},
	}

	EmptyMenu = &tele.ReplyMarkup{RemoveKeyboard: true}
	CreateBtn = &BaseMenu.ReplyKeyboard[0][0]
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

	var userList = make(map[int64]*UserContext, 100)
	var handler = &FirstNameHandler{}
	handler.SetNext(&SecondNameHandler{}).
		SetNext(&DateHandler{}).
		SetNext(&FilterHandler{}).
		SetNext(&GeneratorHandler{})

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hey there! ☺️ We're Imagen, a bot that creates beautiful, "+
			"personalized images for couples. Whether you're celebrating your first "+
			"date or your 50th anniversary, we've got you covered.\n\nJust press the "+
			"'create' button and we'll generate two wallpapers that form a single, "+
			"unique picture just for you and your significant other. It's a fun and "+
			"easy way to celebrate your love and create something special together. 💜\n\n", BaseMenu)
	})

	b.Handle(CreateBtn, func(c tele.Context) error {
		u := &UserContext{State: FirstNameInput}
		userList[c.Sender().ID] = u
		return handler.Prologue(u, c)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		id := c.Sender().ID
		u, exists := userList[id]
		if !exists || handler.Handle(u, c) != nil {
			return c.Send("To start just press create button 👇️", BaseMenu)
		}

		if u.State == ImageOutput {
			delete(userList, id)
		}
		return nil
	})

	b.Start()
}
