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
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hasokon/mahjan"
	"github.com/line/line-bot-sdk-go/linebot"
)

type MahjanScore struct {
	person mahjan.Person
	tsumo  bool
	hu     uint
	han    uint
}

func (this MahjanScore) getMahjanScore() string {
	m := mahjan.New()
	return m.Score(this.hu, this.han, this.person, this.tsumo)
}

func (this MahjanScore) String() string {
	p := "親"
	if this.person == mahjan.Child {
		p = "子"
	}

	t := "ロン"
	if this.tsumo {
		t = "ツモ"
	}

	return fmt.Sprintf("%s %s %d符%d翻", p, t, this.hu, this.han)
}

func replyMahjanYaku() string {
	m := mahjan.New()

	yaku := ""
	for _, v := range m.Yakulist {
		yaku = yaku + v.String() + "\n"
	}

	return yaku
}

func reply(bot *linebot.Client, text string, event *linebot.Event) {
	message := ""
	r := regexp.MustCompile(`ンゴ$`)
	switch {
	case text == "334":
		message = "なんでや！阪神関係ないやろ！"
	case r.MatchString(text):
		message = "はえ〜"
	case text == "麻雀の役を教えて":
		message = replyMahjanYaku()
	case text == "麻雀の点数計算して":
		replyParentOrChild(bot, event)
	default:
		return
	}
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
	}
}

func replyParentOrChild(bot *linebot.Client, event *linebot.Event) {
	parentAction := linebot.NewPostbackTemplateAction("親", "parent_or_child,parent", "")
	childAction := linebot.NewPostbackTemplateAction("子", "parent_or_child,child", "")

	template := linebot.NewButtonsTemplate("", "", "親 or 子", parentAction, childAction)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("親 or 子", template)).Do(); err != nil {
		log.Print(err)
	}
}

func replyTsumoOrRon(bot *linebot.Client, event *linebot.Event) {
	tsumoAction := linebot.NewPostbackTemplateAction("ツモ", "tsumo_or_ron,tsumo", "")
	ronAction := linebot.NewPostbackTemplateAction("ロン", "tsumo_or_ron,ron", "")

	template := linebot.NewButtonsTemplate("", "", "ツモ or ロン", tsumoAction, ronAction)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("ツモ or ロン", template)).Do(); err != nil {
		log.Print(err)
	}
}

func replyHu(bot *linebot.Client, event *linebot.Event) {
	action20 := linebot.NewPostbackTemplateAction("20", "hu,20", "")
	action25 := linebot.NewPostbackTemplateAction("25", "hu,25", "")
	action30 := linebot.NewPostbackTemplateAction("30", "hu,30", "")
	action40 := linebot.NewPostbackTemplateAction("40", "hu,40", "")

	template := linebot.NewButtonsTemplate("", "", "何符？", action20, action25, action30, action40)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("何符？", template)).Do(); err != nil {
		log.Print(err)
	}
}

func replyHan(bot *linebot.Client, event *linebot.Event) {
	action1 := linebot.NewPostbackTemplateAction("1", "han,1", "")
	action2 := linebot.NewPostbackTemplateAction("2", "han,2", "")
	action3 := linebot.NewPostbackTemplateAction("3", "han,3", "")
	action4 := linebot.NewPostbackTemplateAction("4", "han,4", "")

	template := linebot.NewButtonsTemplate("", "", "何翻？", action1, action2, action3, action4)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("何翻？", template)).Do(); err != nil {
		log.Print(err)
	}
}

func replyFromImage(bot *linebot.Client, id string, event *linebot.Event) {
	content, err := bot.GetMessageContent(id).Do()
	if err != nil {
		log.Print(err)
	}
	defer content.Content.Close()

	labels, err := FindLabels(content.Content)
	if err != nil {
		log.Print(err)
		return
	}

	messages := make([]linebot.Message, 0)
	for _, v := range labels {
		messages = append(messages, linebot.NewTextMessage(v))
	}

	if _, err := bot.ReplyMessage(event.ReplyToken, messages[:5]...).Do(); err != nil {
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
					reply(bot, message.Text, event)
				case *linebot.ImageMessage:
					replyFromImage(bot, message.ID, event)
				}
			}

			if event.Type == linebot.EventTypePostback {
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
