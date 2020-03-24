package discord

import (
	"encoding/json"
	"fmt"
	"time"
)

// Heartbeat ...
func (c *Client) heartbeat(interval int) {
	defer c.waitGroup.Done()
	heartbeatBody := []byte(fmt.Sprintf(`{
		"op": 1,
		"d": %d
	}`, c.seq))

	for {
		if c.heartbeatAck == true {
			c.heartbeatAck = false
			c.SendEvent(heartbeatBody)
			time.Sleep(time.Duration(interval) * time.Millisecond)
		} else {
			panic("Heartbeat Not Acknowledged")
		}
	}
}

// ExtractInterval ...
func ExtractInterval(body json.RawMessage) (int, error) {
	res := HeartBeatRes{}
	err := json.Unmarshal(body, &res)
	if err != nil {
		return 0, err
	}
	return res.Interval, nil
}
