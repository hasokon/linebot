package main

import (
	"log"

	"github.com/hasokon/mahjan"
	"github.com/line/line-bot-sdk-go/linebot"
)

func replyMahjanYaku() string {
	m := mahjan.New()

	yaku := ""
	for _, v := range m.Yakulist {
		yaku = yaku + v.String() + "\n"
	}

	return yaku
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
