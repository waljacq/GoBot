package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type event struct {
	T  string          `json:"t"`
	S  int             `json:"s"`
	OP int             `json:"op"`
	D  json.RawMessage `json:"d"`
}

type discordURL struct {
	URL string `json:"url"`
}

// Initialize ...
func Initialize() {
	if os.Getenv("BOT_AUTH") == "" {
		fmt.Println("Please set your 'BOT_AUTH' environment Variable")
		os.Exit(1)
	}

	//token := os.Getenv("BOT_AUTH")
	header := http.Header{}
	url, err := gateway()
	if err != nil {
		panic(err)
	}
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		panic(err)
	}

	conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})

	msgtype, msg, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	var newEvent event
	json.Unmarshal(msg, &newEvent)
	fmt.Printf("\n%+v\n", newEvent)

	fmt.Printf("\n%d - %s\n", msgtype, string(msg))
}

func gateway() (string, error) {
	resp, err := http.Get("https://discordapp.com/api/v6/gateway?v=6&encoding=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	url := discordURL{}
	err = json.Unmarshal(body, &url)
	if err != nil {
		return "", err
	}
	return url.URL, nil
}

// https://discordapp.com/api/v6/gateway?v=6&encoding=json
