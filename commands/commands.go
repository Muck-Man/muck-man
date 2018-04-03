package commands

import (
	"github.com/bwmarrin/discordgo"
)

type command interface {
	exec(s *discordgo.Session, m *discordgo.Message)
}

var commands = make(map[string]command)
