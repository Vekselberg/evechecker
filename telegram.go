package main

import (
	"strconv"

	"gopkg.in/telegram-bot-api.v4"

	"log"
)

func findid() {
	_, _, _, _, botid, _ := config()

	bot, err := tgbotapi.NewBotAPI(botid)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Println("Please send any message to bot to obtain your ID.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	//text := "sometext"
	//msg := tgbotapi.NewMessage(119833076, text)
	//msg.ReplyToMessageID = update.Message.MessageID

	//bot.Send(msg)

	for update := range updates {
		if update.Message == nil {
			break
		}

		//fmt.Println(update)

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		userid := strconv.Itoa(update.Message.From.ID)

		text := "Your userID: " + userid + "\n\nSet it in your eveapi.cfg file."
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
		break
	}

	log.Println("END")

}

func achtung(what string) {
	_, _, _, _, botid, userid := config()
	bot, err := tgbotapi.NewBotAPI(botid)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	text := what + " Under attack!"
	msg := tgbotapi.NewMessage(userid, text)
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
