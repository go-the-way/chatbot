package messagehandlers

import "github.com/go-the-way/chatbot/sdk"

func chatGPT(param string) Body {
	return Body{
		At:      At{IsAtAll: true},
		MsgType: "text",
		Text:    Text{sdk.Completion(param)},
	}
}
