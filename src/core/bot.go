package core

import (
	"fmt"
	commands "main/src"
	events "main/src/Events"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Bot    BotConfig
	Client ClientConfig
}

type BotConfig struct {
	DiscordToken string `toml:"discord_token"`
}

type ClientConfig struct {
	ClientID string `toml:"client_id"`
}

func BotConnect() {
	tomlFilePath := "src/files/config.toml"

	file, err := os.Open(tomlFilePath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier TOML:", err)
		return
	}
	defer file.Close()

	var config Config

	// Chargez les données depuis le fichier TOML en décodant dans la structure
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Erreur lors du chargement du fichier TOML:", err)
		return
	}

	// connection avec le token
	dg, err := discordgo.New("Bot " + config.Bot.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the events for the bot.
	dg.AddHandler(commands.HandleCommand)
	dg.AddHandler(events.MessageCreate)
	dg.AddHandler(events.MessageDelete)
	dg.AddHandler(events.Ready)

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
