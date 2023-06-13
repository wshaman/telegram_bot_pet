package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"./exchange"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	data, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatal(err)
	}
	token := string(data)

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-messages
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/currency") {
			currencyOptions := []string{"AUD", "USD", "RUB"}
			var rows []tgbotapi.KeyboardButton
			for _, currency := range currencyOptions {
				button := tgbotapi.NewKeyboardButton(currency)
				rows = append(rows, button)
			}
			replyMarkup := tgbotapi.NewReplyKeyboard(rows)
			replyMarkup.OneTimeKeyboard = true
			response := "Please select a currency:"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyMarkup = replyMarkup
			bot.Send(msg)
			continue

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		fmt.Println(msg.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		currency := exchange.Exchange(msg.Text)

		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.UserName+"!"+"\n"+currency.Title+" "+currency.Current)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
		continue
	}
}
