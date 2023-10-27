package core

import (
	"fmt"
	commands "main/src"
	events "main/src/Events"
	"main/src/core/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func BotConnect() {

	dg, err := discordgo.New("Bot " + config.Main.Bot.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the events for the bot.
	dg.AddHandler(commands.HandleCommand)

	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.MessageDelete)

	dg.AddHandler(events.Ready)

	dg.AddHandler(events.ChannelCreate)

	dg.AddHandler(events.GuildBanAdd)
	dg.AddHandler(events.GuildBanRemove)

	dg.AddHandler(events.GuildCreate)
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
