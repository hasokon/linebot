package main

import (
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func createCallbackHandler(bot *linebot.Client) func(w http.ResponseWriter, req *http.Request) {

	ms := MahjanScore{hu: 40, han: 3}

	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if REPLY_TEXT_MESSAGE {
						reply(bot, message.Text, event)
					}
				case *linebot.ImageMessage:
					if REPLY_IMAGE {
						replyFromImage(bot, message.ID, event)
					}
				}
			}
			if event.Type == linebot.EventTypePostback && MAHJAN_SCORE {
				postback := event.Postback
				datas := strings.Split(postback.Data, ",")
				replyMahjan(ms, datas[0], datas[1], bot, event)
			}
		}
	}
}
