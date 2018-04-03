package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	commands["ping"] = new(pingCommand)
	commands["panic"] = new(panicCommand)
}

type pingCommand struct{}

func (p *pingCommand) exec(s *discordgo.Session, m *discordgo.Message) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "pong!"); err != nil {
		report(err)
	}
}

type panicCommand struct{}

func (p *panicCommand) exec(s *discordgo.Session, m *discordgo.Message) {
	panic("2+2 is 4")
}
