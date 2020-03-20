package discord

import "fmt"

func (c *Client) listening() {
	defer c.waitGroup.Done()
	for {
		newEvent := c.ReadEvent()
		// fmt.Printf("\n\nReceived new event: %v \n\n", newEvent)
		c.handleEvent(newEvent)
		PrintEvent(newEvent)
	}
}

func (c *Client) handleEvent(eve Event) {

	// We can fill out these OP codes as we find things we need.
	switch eve.OP {
	case 0:
		fmt.Println("\n\nOP code 0 received")
	case 1:
		fmt.Println("\n\nOP code 1 received")
	case 2:
		fmt.Println("\n\nOP code 2 received")
	case 3:
		fmt.Println("\n\nOP code 3 received")
	case 11:
		fmt.Println("\n\nOP code 11 received")
		c.heartbeatAck = true
	default:
		fmt.Println("\n\nAn unsupported OP code was received")
	}
}
