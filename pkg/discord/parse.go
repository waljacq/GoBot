package discord

import (
	"encoding/json"
	"fmt"
	"go-bot/pkg/commands"
)

func parseDispatch(body json.RawMessage, tag string) {
	if tag == "MESSAGE_CREATE" {
		var msg commands.Message
		json.Unmarshal([]byte(body), &msg)
		fmt.Printf("\nMessage Object: %+v\n", msg)

		if msg.Content[0] != '!' {
			return
		}

		// Chop of the '!'
		msg.Content = msg.Content[1:]

		// potentially spin up a go routine in order to not block the main thread on serving a command
		commands.Parse(msg)
	}
}
