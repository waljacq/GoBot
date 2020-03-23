package discord

import (
	"encoding/json"
	"fmt"
)

// ReadEvent ...
func (c *Client) ReadEvent() Event {
	_, msg, err := c.conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	var newEvent Event
	json.Unmarshal(msg, &newEvent)
	return newEvent
}

// SendEvent ...
func (c *Client) SendEvent(msg []byte) {
	c.connLock.Lock()
	err := c.conn.WriteMessage(1, msg)
	c.connLock.Unlock()
	if err != nil {
		panic(err)
	}
}

// PrintEvent ...
func PrintEvent(eve Event) {
	fmt.Printf(`
	OP: %d
	T: %s
	S: %d
	D: %s
	`, eve.OP, eve.T, eve.S, eve.D)
}

func (c *Client) handleEvent(eve Event) {

	if eve.S != 0 {
		c.seq = eve.S
	}

	// We can fill out these OP codes as we find things we need.
	switch eve.OP {
	case 0:
		fmt.Println("\n\nOP code 0 received")
		parseDispatch(eve.D, eve.T)
	case 1:
		fmt.Println("\n\nOP code 1 received")
	case 2:
		fmt.Println("\n\nOP code 2 received")
	case 3:
		fmt.Println("\n\nOP code 3 received")
	case 4:
		fmt.Println("\n\nOP code 4 received")
	case 5:
		fmt.Println("\n\nOP code 5 received")
	case 6:
		fmt.Println("\n\nOP code 6 received")
	case 7:
		fmt.Println("\n\nOP code 7 received")
	case 8:
		fmt.Println("\n\nOP code 8 received")
	case 9:
		fmt.Println("\n\nOP code 9 received")
	case 10:
		fmt.Println("\n\nOP code 10 received")
	case 11:
		fmt.Println("\n\nOP code 11 received")
		c.heartbeatAck = true
	default:
		fmt.Println("\n\nAn unsupported OP code was received")
	}
}
