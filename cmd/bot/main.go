package main

import (
	"log"

	"github.com/anuarsabitovich/telegram_bot_pet/app/bot"
	"github.com/anuarsabitovich/telegram_bot_pet/app/config"
)

func main() { //@todo
	//token := os.Getenv("TG_APP_TOKEN")
	//data, err := ioutil.ReadFile("token.txt") //@todo
	//if err != nil {
	//	log.Fatal(err)
	//}
	//token := string(data)
	cfg := config.Load()
	if err := bot.Run(cfg.Token); err != nil {
		//send err to external logger
		log.Fatal(err)
	}
}
