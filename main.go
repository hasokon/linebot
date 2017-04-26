package main

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	// REPLY_TEXT_MESSAGE is enable reply from text message
	REPLY_TEXT_MESSAGE = true
	// REPLY_IMAGE is enable reply from image message
	REPLY_IMAGE = true
	// REPLY_MOUKOBEN is enable reply from moukoben
	REPLY_MOUKOBEN = true
	// MAHJAN_SCORE is enable mahjan score function
	MAHJAN_SCORE = false
	// MAHJAN_YAKU is enable printing mahjan yaku list
	MAHJAN_YAKU = true
	// CHECK_MESHITERO is enable meshitero checker
	CHECK_MESHITERO = true
	// CHECK_SAFETY is enable image safety checker
	CHECK_SAFETY = true
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", createCallbackHandler(bot))

	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
