package commands

import (
	"bytes"
	"fmt"
	"go-bot/pkg/metadata"
	"net/http"
	"os"
	"time"
)

// CMDS ...
var CMDS = map[string]func(msg Message){
	"ping":   ping,
	"uptime": uptime,
}

func ping(msg Message) {
	data := []byte(`{"content": "Pong!"}`)
	client := &http.Client{}
	url := fmt.Sprintf("https://discordapp.com/api/v6/channels/%s/messages", msg.ChannelID)
	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bot %s", os.Getenv("BOT_AUTH")))
	request.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func uptime(msg Message) {
	currentTime := time.Now()
	elapsedTime := currentTime.Sub(metadata.StartTime)
	contentString := fmt.Sprintf(`{"content": "%v"}`, elapsedTime)

	data := []byte(contentString)
	client := &http.Client{}
	url := fmt.Sprintf("https://discordapp.com/api/v6/channels/%s/messages", msg.ChannelID)
	body := bytes.NewBuffer(data)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bot %s", os.Getenv("BOT_AUTH")))
	request.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
