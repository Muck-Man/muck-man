package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var mention = ""

// Register the commands plugin
func Register(discord *discordgo.Session) {
	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)
}

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	mention = e.User.Mention()
}

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author != nil && e.Author.Bot {
		return
	}
	cmd := ""
	if strings.HasPrefix(e.Content, ".m ") {
		cmd = e.Content[3:]
	} else if strings.HasPrefix(e.Content, mention) {
		cmd = e.Content[len(mention)+1:]
	} else {
		return
	}

	parts := strings.Split(cmd, " ")
	if len(parts) < 0 {
		return
	}
	command, set := commands[parts[0]]
	if !set {
		return
	}
	command.exec(s, e.Message)
}
