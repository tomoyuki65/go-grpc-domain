package chat

import (
	"fmt"
	"strings"
	"time"
)

type Chat struct {
	InputText string `json:"input_text"`
}

func NewChat(inputText string) *Chat {
	return &Chat{
		InputText: inputText,
	}
}

func (c *Chat) TextToUpper() string {
	return strings.ToUpper(c.InputText)
}

func (c *Chat) TextToLower() string {
	return strings.ToLower(c.InputText)
}

func (c *Chat) TextAddTimeNow() string {
	return fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), c.InputText)
}
