package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	commands["ping"] = new(pingCommand)
}

type pingCommand struct{}

func (p *pingCommand) exec(s *discordgo.Session, m *discordgo.Message) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "pong!"); err != nil {
		report(err)
	}
}
