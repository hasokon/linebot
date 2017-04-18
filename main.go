// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"
	"net/http"
	"os"
//	"regexp"
	"strings"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/hasokon/mahjan"
)

func replyMahjanScore(text string) string {
	m := mahjan.New()

	person := mahjan.Parent
	if text[1] == 'c' {
		person = mahjan.Child
	}

	tsumo := true
	if text[2] == 'r' {
		tsumo = false
	}

	text = text[3:]
	nums := strings.Split(text, ",")
	hu,_ := strconv.Atoi(nums[0])
	han,_ := strconv.Atoi(nums[1])

	return m.Score(uint(hu), uint(han), person, tsumo)
}

func replyMahjanYaku() string {
	m := mahjan.New()

	yaku := ""
	for _, v := range m.Yakulist {
		yaku = yaku + v.String() + "\n"
	}

	return yaku
}
/*
func reply(bot *linebot.Client, text string, event *linebot.Event) {
	message := ""
	r := regexp.MustCompile(`ンゴ$`)
	mahjan := regexp.MustCompile(`^m[pc][tr][0-9]*,[0-9]`)
	switch {
	case text == "334":
		message = "なんでや！阪神関係ないやろ！"
	case r.MatchString(text):
		message = "はえ〜"
	case mahjan.MatchString(text):
		message = replyMahjanScore(text)
	case text == "麻雀の役を教えて":
		message = replyMahjanYaku()
	default:
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
	}
}
*/

func reply(bot *linebot.Client, text string, event *linebot.Event) {
	parentAction := linebot.NewPostbackTemplateAction("Parent", "parent data", "")
	childAction := linebot.NewPostbackTemplateAction("Child", "child data", "")

	template := linebot.NewButtonsTemplate("", "", "Who are you?", parentAction, childAction)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("test",template)).Do(); err != nil {
		log.Print(err)
	}
}

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
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
					reply(bot, message.Text, event)
				}
			}

			if event.Type == linebot.EventTypePostback {
				postback := event.Postback
				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(postback.Data)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
