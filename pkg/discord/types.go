package discord

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

// Event ...
type Event struct {
	T  string          `json:"t"`
	S  int             `json:"s"`
	OP int             `json:"op"`
	D  json.RawMessage `json:"d"`
}

// URL ...
type URL struct {
	URL string `json:"url"`
}

// HeartBeatRes ...
type HeartBeatRes struct {
	Interval int `json:"heartbeat_interval"`
}

// Client represents the state data and connection to discord.
type Client struct {
	conn         *websocket.Conn
	waitGroup    *sync.WaitGroup
	connLock     *sync.Mutex
	heartbeatAck bool
	seq          int
}
