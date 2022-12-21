package bot

import "strings"

type PromptState int32

const (
	FirstNameInput PromptState = iota
	SecondNameInput
	DateInput
)

type User struct {
	FName string
	SName string
	Date  string
	State PromptState
}

func (u User) String() string {
	var sb strings.Builder
	sb.WriteString(u.FName)
	sb.WriteString(" ")
	sb.WriteString(u.SName)
	sb.WriteString(" ")
	sb.WriteString(u.Date)
	sb.WriteString(" ")
	return sb.String()
}
