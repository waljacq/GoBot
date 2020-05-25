package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Initialize ...
func Initialize() {
	if os.Getenv("BOT_AUTH") == "" {
		fmt.Println("Please set your 'BOT_AUTH' environment Variable")
		os.Exit(1)
	}

	header := http.Header{}
	url, err := gateway()
	if err != nil {
		panic(err)
	}
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	waitGroup := sync.WaitGroup{}
	connLock := sync.Mutex{}
	c := Client{conn, &waitGroup, &connLock, time.Now(), true, 0}
	if err != nil {
		panic(err)
	}

	conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})
	defer conn.Close()

	newEvent := c.ReadEvent()
	interval, err := ExtractInterval(newEvent.D)
	if err != nil {
		panic(err)
	}

	identifyMsg := []byte(fmt.Sprintf(`{
		"op": 2,
		"d": {
			"token": "%s",
			"properties": {
				"$os": "linux",
				"$browser": "disco",
				"$device": "disco"
			}
		}
	}`, os.Getenv("BOT_AUTH")))

	c.SendEvent(identifyMsg)

	c.waitGroup.Add(2)
	// Start go routine to send heartbeat message every X ms
	// may need to set up an additional helper function for all message sends to go through for mutex locking
	go c.heartbeat(interval)
	go c.listening()

	c.waitGroup.Wait()
}

func gateway() (string, error) {
	resp, err := http.Get("https://discordapp.com/api/v6/gateway?v=6&encoding=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	url := URL{}
	err = json.Unmarshal(body, &url)
	if err != nil {
		return "", err
	}
	return url.URL, nil
}

// https://discordapp.com/api/v6/gateway?v=6&encoding=json
