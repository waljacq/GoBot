package main

import (
	"go-bot/pkg/discord"
	"os"
	"os/signal"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	discord.Initialize()
}
