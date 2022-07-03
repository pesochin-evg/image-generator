package main

import (
	"net/http"
	"os"

	"github.com/Antipascal/image-generator/pkg/bot"
)

// func MainHandler(resp http.ResponseWriter, _ *http.Request) {
// 	resp.Write([]byte("Just for heroku"))
// }

func main() {
	// http.HandleFunc("/", MainHandler)
	// go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	bot.Start()
}
