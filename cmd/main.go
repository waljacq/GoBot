package main

import (
	"fmt"
	"go-bot/pkg/discord"
	"go-bot/pkg/metadata"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT)
	go func() {
		<-interrupt
		fmt.Println("\nBot Terminated By User")
		os.Exit(0)
	}()
	discord.Initialize()
}

func init() {
	metadata.StartTime = time.Now()
}
