package commands

import (
	"github.com/bwmarrin/discordgo"
)

var commandHandlers = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))
var registeredCommands = make([]*discordgo.ApplicationCommand, 0)

// AddCommand ajoute une commande au gestionnaire
func AddCommand(command *discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	print("add command tqt ")
	commandHandlers[command.Name] = handler
	registeredCommands = append(registeredCommands, command)
}

// RegisterCommands enregistre toutes les commandes dans Discord
func RegisterCommands(s *discordgo.Session, botID, guildID string) {
	println("test de registercommands")
	println(registeredCommands)

	for _, command := range registeredCommands {
		_, err := s.ApplicationCommandCreate(botID, guildID, command)
		if err != nil {
			panic(err)
		}
	}
}

// HandleCommand gère les interactions de commande
func HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, exists := commandHandlers[i.ApplicationCommandData().Name]; exists {
		handler(s, i)
	}
}
