package messagehandlers

type Body struct {
	At      At     `json:"at"`
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type At struct {
	IsAtAll bool `json:"isAtAll"`
}

type Text struct {
	Content string `json:"content"`
}
