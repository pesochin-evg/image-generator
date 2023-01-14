package bot

import (
	"errors"
	tele "gopkg.in/telebot.v3"
)

type Handler interface {
	Handle(c *UserContext, t tele.Context) error
	Prologue(c *UserContext, t tele.Context) error
	SetNext(h Handler) Handler
}

// FirstNameHandler is a handler for the first name input
type FirstNameHandler struct {
	next Handler
}

func (h *FirstNameHandler) Handle(c *UserContext, t tele.Context) error {
	if c.State == FirstNameInput {
		c.FName = t.Text()
		return h.next.Prologue(c, t)
	}

	return h.next.Handle(c, t)
}

func (h *FirstNameHandler) Prologue(c *UserContext, t tele.Context) error {
	c.State = FirstNameInput
	return t.Send("1️⃣ To create your personalized image, we need to know your name. "+
		"\nPlease enter your name below:", EmptyMenu)
}

func (h *FirstNameHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// SecondNameHandler is a handler for the second name input
type SecondNameHandler struct {
	next Handler
}

func (h *SecondNameHandler) Handle(c *UserContext, t tele.Context) error {
	if c.State == SecondNameInput {
		c.SName = t.Text()
		return h.next.Prologue(c, t)
	}

	return h.next.Handle(c, t)
}

func (h *SecondNameHandler) Prologue(c *UserContext, t tele.Context) error {
	c.State = SecondNameInput
	return t.Send("2️⃣ Now, we need to know the name of your significant " +
		"other. \nPlease enter their name below.")
}

func (h *SecondNameHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// DateHandler is a handler for the date input
type DateHandler struct {
	next Handler
}

func (h *DateHandler) Handle(c *UserContext, t tele.Context) error {
	if c.State == DateInput {
		c.Date = t.Text()
		return h.next.Prologue(c, t)
	}

	return h.next.Handle(c, t)
}

func (h *DateHandler) Prologue(c *UserContext, t tele.Context) error {
	c.State = DateInput
	return t.Send("3️⃣ Great! The final piece of information we need is the " +
		"date you started your relationship. \nPlease enter the date in the " +
		"format MM/DD/YYYY:")
}

func (h *DateHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// FilterHandler is a handler for the filter choice
type FilterHandler struct {
	next Handler
}

func (h *FilterHandler) Handle(c *UserContext, t tele.Context) error {
	if c.State == FilterInput {
		c.Filter = t.Text()
		return h.next.Prologue(c, t)
	}

	return h.next.Handle(c, t)
}

func (h *FilterHandler) Prologue(c *UserContext, t tele.Context) error {
	c.State = FilterInput
	return t.Send("4️⃣ Finally, we need to know what filter you want to "+
		"apply to your image. \nPlease select one of the options below:",
		FilterMenu)
}

func (h *FilterHandler) SetNext(next Handler) Handler {
	h.next = next
	return next
}

// GeneratorHandler is a handler for the image generation
type GeneratorHandler struct{}

func (h *GeneratorHandler) Handle(c *UserContext, t tele.Context) error {
	if c.State == ImageOutput {
		return sendImages(c, t)
	}
	return errors.New("unknown command")
}

func (h *GeneratorHandler) Prologue(c *UserContext, t tele.Context) error {
	c.State = ImageOutput
	return h.Handle(c, t)
}

func (h *GeneratorHandler) SetNext(_ Handler) Handler {
	return h
}
