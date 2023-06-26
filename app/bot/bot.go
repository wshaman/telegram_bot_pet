package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/anuarsabitovich/telegram_bot_pet/app/exchange"
)

func connect(token string) (*tgbotapi.BotAPI, error) {
	b, err := tgbotapi.NewBotAPI(token) //@todo Naming

	if err != nil {
		return nil, fmt.Errorf("connecting to bot: %w", err)
	}

	b.Debug = true //@todo : Do not set fields

	// log.Printf("Authorized on account %s", bot.Self.UserName)

	return b, nil
}

func getUpdatesChan(b *tgbotapi.BotAPI, offset, to int) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(offset)
	u.Timeout = to                      //@todo : Do not set fields
	updates, err := b.GetUpdatesChan(u) //@todo : err handling. Extract.
	if err != nil {
		return nil, fmt.Errorf("getting updates chan: %w", err)
	}
	return updates, nil
}

func isMessageProcess(text string) bool {
	return strings.HasPrefix(text, "/currency")
}

func Run(token string) error {
	b, err := connect(token)
	if err != nil {
		return fmt.Errorf("run bot: %w", err)
	}

	updates, err := getUpdatesChan(b, 0, 60)
	for update := range updates {
		if update.Message == nil { // ignore non-messages
			continue
		}

		if isMessageProcess(update.Message.Text) {
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
			b.Send(msg)
			continue

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		fmt.Println(msg.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		currency := exchange.GetCurrentRate(msg.Text)

		//@todo : template
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.UserName+"!"+"\n"+currency.Title+" "+currency.Current)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := b.Send(msg); err != nil {
			fmt.Println(err.Error())
		} //@todo : err handling
		continue
	}
	return nil
}
