package bot

import (
	"strings"
)

type PromptState int32

const (
	FirstNameInput PromptState = iota
	SecondNameInput
	DateInput
	FilterInput
	ImageOutput
)

type UserContext struct {
	FName  string
	SName  string
	Date   string
	State  PromptState
	Filter string
}

func (u *UserContext) String() string {
	var sb strings.Builder
	sb.WriteString(u.FName)
	sb.WriteString(" ")
	sb.WriteString(u.SName)
	sb.WriteString(" ")
	sb.WriteString(u.Date)
	sb.WriteString(" ")
	sb.WriteString(u.Filter)
	sb.WriteString(" ")
	return sb.String()
}
