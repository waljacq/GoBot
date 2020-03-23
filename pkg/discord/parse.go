package discord

import (
	"encoding/json"
	"fmt"
)

func parseDispatch(body json.RawMessage, tag string) {
	if tag == "MESSAGE_CREATE" {
		var msg Message
		json.Unmarshal([]byte(body), &msg)
		fmt.Printf("\nMessage Object: %+v\n", msg)
	}
}
