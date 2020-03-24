package discord

func (c *Client) listening() {
	defer c.waitGroup.Done()
	for {
		newEvent := c.ReadEvent()
		PrintEvent(newEvent)
		c.handleEvent(newEvent)
	}
}
