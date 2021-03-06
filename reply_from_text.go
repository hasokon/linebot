package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func reply(bot *linebot.Client, text string, event *linebot.Event) {
	message := ""
	resultMoukoben, isMoukoben := CheckMoukoben(text)
	resultProg, isProgName := CheckProgramingLangageName(text)
	switch {
	case isMoukoben && REPLY_MOUKOBEN:
		message = resultMoukoben
	case isProgName, REPLY_PROGNAME:
		message = resultProg
	case text == "麻雀の役を教えて" && MAHJAN_YAKU:
		message = replyMahjanYaku()
	case text == "麻雀の点数計算して" && MAHJAN_SCORE:
		replyParentOrChild(bot, event)
		return
	default:
		return
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
	}
}
