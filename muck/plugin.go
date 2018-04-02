package muck

import (
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

var api *muckAPI

// Register the Muck plugin with a discordgo session
func Register(discord *discordgo.Session) {
	api = &muckAPI{
		host:   os.Getenv("MUCK_API_HOST"),
		token:  os.Getenv("MUCK_API_TOKEN"),
		client: http.DefaultClient,
	}

	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onMessageUpdate)
}

func guildID(s *discordgo.Session, channelID string) string {
	c, err := s.State.Channel(channelID)
	if err != nil {
		report(err)
	}
	if c == nil {
		return "null"
	}
	if c.GuildID == "" {
		return "null"
	}
	return c.GuildID
}

func shouldIgnore(m *discordgo.Message) bool {
	if m.Author == nil {
		return true
	}
	if m.Content == "" {
		return true
	}
	return false
}

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if shouldIgnore(e.Message) {
		return
	}

	guild := guildID(s, e.ChannelID)
	data := newMessage(e.Message, guild, false)
	api.sendMessage(data)
}

func onMessageUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	if shouldIgnore(e.Message) {
		return
	}

	guild := guildID(s, e.ChannelID)
	data := newMessage(e.Message, guild, true)
	api.sendMessage(data)
}
