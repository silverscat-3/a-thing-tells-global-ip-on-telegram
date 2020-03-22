package main

import (
	"io/ioutil"
	"log"
	"net/http"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/integrii/flaggy"
)

var (
	version  string = "Schwimmwagen"
	revision string = "Typ 166"
)

func main() {
	var token string
	flaggy.String(&token, "t", "token", "Telegram token")
	flaggy.Parse()

	if "" == token {
		log.Fatalln("Token is required.")
	}

	bot, err := tg.NewBotAPI(token)
	if nil != err {
		log.Fatalln(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tg.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "help":
			msg.Text = "/showでグローバルIP表示するで"

		case "show":
			ip, err := getGlobalIP()
			if nil != err {
				log.Fatalln(err)
			}
			msg.Text = ip

		default:
			msg.Text = "そんなコマンド知らん"
		}

		if _, err := bot.Send(msg); nil != err {
			log.Fatalln(err)
		}
	}
}

func getGlobalIP() (string, error) {
	resp, err := http.Get("http://inet-ip.info/ip")
	if nil != err {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return "", err
	}

	return string(body), nil
}
