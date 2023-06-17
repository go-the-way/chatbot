package messagehandlers

import (
	"strings"

	"github.com/go-the-way/chatbot/sdk"
)

func Call(content string) { sdk.SendMessage(chatGPT(strings.TrimSpace(content))) }
