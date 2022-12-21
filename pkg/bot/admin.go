package bot

import (
	tele "gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

type Admins struct {
	users []tele.User
}

var instance *Admins = nil

func (a Admins) Notify(what interface{}, c tele.Context) error {
	for _, user := range a.users {
		b := c.Bot()
		_, err := b.Send(&user, what)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAdmins() Admins {
	if instance == nil {
		u := make([]tele.User, 1)
		u[0] = tele.User{
			ID:           541204191,
			FirstName:    "Evgeniy",
			LastName:     "Pesochin",
			Username:     "BRTL16",
			LanguageCode: "en",
		}
		instance = &Admins{users: u}
	}
	return *instance
}

func GetUserString(c tele.Context) string {
	var sb strings.Builder
	sb.WriteString("* ")
	sb.WriteString(c.Sender().FirstName)
	sb.WriteString(" ")
	sb.WriteString(c.Sender().LastName)
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatInt(c.Sender().ID, 10))
	sb.WriteString(" https://t.me/@")
	sb.WriteString(c.Sender().Username)
	return sb.String()
}
