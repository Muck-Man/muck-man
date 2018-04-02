package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	fmt.Println("startup.")
	if err = discord.Open(); err != nil {
		panic(err)
	}

	// block until asked to close
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("shutdown.")
	discord.Close()
}

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	fmt.Printf("%s is ready\n", e.User.String())
}

func onMessageCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	fmt.Printf("(%v) %q\n", e.ChannelID, e.Content)
}
