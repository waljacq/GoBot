package main

import (
	"go-bot/pkg/discord"
	"fmt"
	"os"
)

func main() {
	if os.Getenv("BOT_AUTH") != "" {
		fmt.Println("Please set your 'BOT_AUTH' environment Variable")
		os.Exit(1)
	}
	discord.New()
}