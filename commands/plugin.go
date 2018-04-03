package commands

import (
	"fmt"
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
	mention = fmt.Sprintf("%s ", e.User.Mention())
}

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author != nil && e.Author.Bot {
		return
	}
	cmd := ""
	if strings.HasPrefix(e.Content, ".m ") {
		cmd = e.Content[3:]
	} else if strings.HasPrefix(e.Content, ".muck ") {
		cmd = e.Content[6:]
	} else if strings.HasPrefix(e.Content, mention) {
		cmd = e.Content[len(mention):]
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

	defer func() {
		if r := recover(); r != nil {
			println("(cmd panic!)", r)
		}
	}()
	command.exec(s, e.Message)
}
