package discord

import (
	"encoding/json"
	"sync"
	"time"

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

// Message ...
type Message struct {
	MsgType   int       `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Mentions  []User    `json:"mentions"`
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	ChannelID string    `json:"channel_id"`
	Author    User      `json:"author"`
	GuildID   string    `json:"guild_id"`
}

// User ...
type User struct {
	Username      string `json:"username"`
	ID            string `json:"id"`
	Discriminator string `json:"discriminator"`
}

// Client represents the state data and connection to discord.
type Client struct {
	conn         *websocket.Conn
	waitGroup    *sync.WaitGroup
	connLock     *sync.Mutex
	heartbeatAck bool
	seq          int
}
