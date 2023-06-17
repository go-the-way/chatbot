package messagehandlers

type HandlerFunc func(param string) Body

type MessageHandler interface {
	Cmd() string
	Alias() []string
	Description() string
	Func() HandlerFunc
}
