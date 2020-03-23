package discord

func (c *Client) listening() {
	defer c.waitGroup.Done()
	for {
		newEvent := c.ReadEvent()
		// fmt.Printf("\n\nReceived new event: %v \n\n", newEvent)
		c.handleEvent(newEvent)
		PrintEvent(newEvent)
	}
}
