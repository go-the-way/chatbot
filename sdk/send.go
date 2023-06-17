package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-the-way/chatbot/pkg"
)

var (
	dingTalkRobotBotToken = pkg.BindEnv("DING_TALK_ROBOT_BOT_TOKEN", "")
	robotSendURL          = "https://oapi.dingtalk.com/robot/send?access_token=" + dingTalkRobotBotToken
)

func SendMessage(msg any) {
	buf, _ := json.Marshal(&msg)
	resp, err := http.Post(robotSendURL, "application/json", bytes.NewReader(buf))
	fmt.Println(err)
	buf, err = io.ReadAll(resp.Body)
	fmt.Println(err)
	fmt.Println(string(buf))
	fmt.Println(resp.StatusCode)
}
