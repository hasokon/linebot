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
	"strconv"
	"strings"

	"github.com/hasokon/mahjan"
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

	ms := MahjanScore{hu: 40, han: 3}

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
				switch datas[0] {
				case "parent_or_child":
					switch datas[1] {
					case "parent":
						ms.person = mahjan.Parent
					case "child":
						ms.person = mahjan.Child
					}
					replyTsumoOrRon(bot, event)
				case "tsumo_or_ron":
					switch datas[1] {
					case "tsumo":
						ms.tsumo = true
					case "ron":
						ms.tsumo = false
					}
					replyHu(bot, event)
				case "hu":
					hu, _ := strconv.Atoi(datas[1])
					ms.hu = uint(hu)
					replyHan(bot, event)
				case "han":
					han, _ := strconv.Atoi(datas[1])
					ms.han = uint(han)
					yaku := linebot.NewTextMessage(ms.String())
					score := linebot.NewTextMessage(ms.getMahjanScore())
					if _, err := bot.ReplyMessage(event.ReplyToken, yaku, score).Do(); err != nil {
						log.Print(err)
					}
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
