package main

import (
	"fmt"
	"net/http"

	"github.com/go-the-way/chatbot/pkg"

	_ "github.com/go-the-way/chatbot/handlers"
)

func main() {
	fmt.Println("chat bot started")
	fmt.Println(http.ListenAndServe(pkg.BindEnv("SERVER_ADDR", ":50000"), nil))
}
