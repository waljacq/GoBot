package commands

import "time"

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
