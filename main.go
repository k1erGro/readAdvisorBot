package main

import (
	"flag"
	"log"
	Client "tg_bot/clients/telegram"
	"tg_bot/consumer/event_consumer"
	"tg_bot/events/telegram"
	"tg_bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	tgClient := Client.New(tgBotHost, mustToken())

	eventsProcessor := telegram.New(tgClient, files.New(storagePath))

	log.Print("Service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
