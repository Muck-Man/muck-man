package commands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	commands["help"] = new(helpCommand)
}

const helpFormat = `
**muck man**

<https://muck.gg>
`

type helpCommand struct{}

func (h *helpCommand) exec(s *discordgo.Session, m *discordgo.Message) {

	if _, err := s.ChannelMessageSend(m.ChannelID, helpFormat); err != nil {
		report(err)
	}
}
