package main

import (
	"Telegramm-botTZ/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	//token = flags.Get(token)
	tgClient := telegram.New(tgBotHost, mustToken())

	//tgClient = telegram.New(token)

	//fetcher = fetcher.New(tgClient)

	//processor = processor.New(tgClient)

	//consumer.Start(fetcher, porocessor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for accass to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
