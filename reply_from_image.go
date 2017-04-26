package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func checkMeshitero(labels []string) bool {
	for _, label := range labels {
		switch {
		case strings.Contains(label, "dish"):
			return true
		case strings.Contains(label, "food"):
			return true
		case strings.Contains(label, "cuisine"):
			return true
		case strings.Contains(label, "meal"):
			return true
		case strings.Contains(label, "drink"):
			return true
		}
	}

	return false
}

func replyMeshitero(bot *linebot.Client, id string, event *linebot.Event) {
	contentForMeshitero, err := bot.GetMessageContent(id).Do()
	if err != nil {
		log.Print(err)
	}
	defer contentForMeshitero.Content.Close()

	labels, err := FindLabels(contentForMeshitero.Content)
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(labels)

	if checkMeshitero(labels) {
		rand.Seed(time.Now().UnixNano())
		message := ""
		switch rand.Intn(4) {
		case 0:
			message = "飯テロを検知したンゴ！"
		case 1:
			message = "お腹が空いてきたンゴね〜"
		case 2:
			message = "飯ヤメて...やめてｸﾚﾒﾝｽ..."
		case 3:
			message = "飯テロ、絶許"
		default:
			return
		}
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
			log.Print(err)
		}
	}
}

func replySafety(bot *linebot.Client, id string, event *linebot.Event) {
	contentForSafeSearch, err := bot.GetMessageContent(id).Do()
	if err != nil {
		log.Print(err)
	}
	defer contentForSafeSearch.Content.Close()
	annotation, err := CheckSafety(contentForSafeSearch.Content)
	if err != nil {
		log.Print(err)
		return
	}
	/*
	   Likelihood_UNKNOWN Likelihood = 0
	   Likelihood_VERY_UNLIKELY Likelihood = 1
	   Likelihood_UNLIKELY Likelihood = 2
	   Likelihood_POSSIBLE Likelihood = 3
	   Likelihood_LIKELY Likelihood = 4
	   Likelihood_VERY_LIKELY Likelihood = 5
	*/
	log.Println(annotation.Adult, annotation.Medical, annotation.Spoof, annotation.Violence)

	switch {
	case annotation.Adult >= 3 || annotation.Medical >= 4 || annotation.Spoof >= 4 || annotation.Violence >= 4:
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("これはいけない")).Do(); err != nil {
			log.Print(err)
		}
		return
	case annotation.Adult >= 2:
		if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("あやしい画像を検知しました")).Do(); err != nil {
			log.Print(err)
		}
	}
}

func replyFromImage(bot *linebot.Client, id string, event *linebot.Event) {
	switch {
	case CHECK_MESHITERO:
		replyMeshitero(bot, id, event)
	case CHECK_SAFETY:
		replySafety(bot, id, event)
	}
}
