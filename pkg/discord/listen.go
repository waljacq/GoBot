package discord

func (c *Client) listening() {
	defer c.waitGroup.Done()
	for {
		newEvent := c.ReadEvent()
		c.handleEvent(newEvent)
		PrintEvent(newEvent)
	}
}
