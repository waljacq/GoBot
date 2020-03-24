package commands

import (
	"strings"
)

// Parse ...
func Parse(msg Message) {
	command := msg.Content
	if strings.Index(command, " ") != -1 {
		cmdParams := command[strings.Index(command, " "):]
		command = command[:strings.Index(command, " ")]
		msg.Content = cmdParams
	}

	CMDS[command](msg)
}
