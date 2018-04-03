package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muck-Man/muck-man/commands"
	"github.com/Muck-Man/muck-man/muck"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	token := fmt.Sprintf("Bot %s", os.Getenv("DISCORD_TOKEN"))

	discord, err := discordgo.New(token)
	if err != nil {
		panic(err)
	}
	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onGuildCreate)
	discord.AddHandler(onGuildDelete)

	commands.Register(discord)
	muck.Register(discord)

	fmt.Println("(bot) startup")
	if err = discord.Open(); err != nil {
		panic(err)
	}

	// block until asked to close
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("(bot) shutting down")
	discord.Close()
	fmt.Println("(bot) bye")
}

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	s.UpdateStatus(0, ".muck help")

	fmt.Printf("(bot) %s is ready, expecting %d guilds\n", e.User.String(), len(e.Guilds))
}
func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Content == "" {
		return
	}
	fmt.Printf("(bot) %v %q\n", e.ChannelID, e.Content)
}
func onGuildCreate(s *discordgo.Session, e *discordgo.GuildCreate) {
	fmt.Printf("(bot-guild) create %s\n", e.Name)
}
func onGuildDelete(s *discordgo.Session, e *discordgo.GuildDelete) {
	fmt.Printf("(bot-guild) delete %s\n", e.Name)
}
