package sdk

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/sashabaranov/go-openai"

	"github.com/go-the-way/chatbot/pkg"
)

var (
	chatGPTHttpProxy      = pkg.BindEnv("CHAT_GPT_HTTP_PROXY", "N") == "T"
	chatGPTHttpProxyAddr  = pkg.BindEnv("CHAT_GPT_HTTP_PROXY_ADDR", "http://localhost:1081")
	chatGPTToken          = pkg.BindEnv("CHAT_GPT_TOKEN", "")
	chatGPTRequestTimeout = pkg.BindEnv("CHAT_GPT_REQUEST_TIMEOUT", "1m")
)

func Completion(content string) (result string) {
	config := openai.DefaultConfig(chatGPTToken)
	config.HTTPClient = getClient()
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{Model: openai.GPT3Dot5Turbo, Messages: []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: content}}})
	if err != nil {
		return "ChatCompletion error: " + err.Error()
	}
	return resp.Choices[0].Message.Content
}

func getClient() *http.Client {
	transport := &http.Transport{}
	if chatGPTHttpProxy {
		if chatGPTHttpProxyAddr != "" {
			proxyURL, err := url.Parse(chatGPTHttpProxyAddr)
			if err == nil {
				transport.Proxy = http.ProxyURL(proxyURL)
			}
		}
	}
	timeout, _ := time.ParseDuration(chatGPTRequestTimeout)
	return &http.Client{Transport: transport, Timeout: timeout}
}
