package commands

import (
	"fmt"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var commandHandlers = make(map[string]func(c bot.Client, i *events.ApplicationCommandInteractionCreate))
var registeredCommands = make([]discord.ApplicationCommandCreate, 0)

// AddCommand ajoute une commande au gestionnaire
func AddCommand(command *discord.SlashCommandCreate, handler func(s bot.Client, i *events.ApplicationCommandInteractionCreate)) {
	commandHandlers[command.Name] = handler
	registeredCommands = append(registeredCommands, command)
}

// RegisterCommands enregistre toutes les commandes dans Discord
func RegisterCommands(client bot.Client, botID, guildID string) {

	fmt.Println(registeredCommands)
	client.Rest().SetGlobalCommands(client.ApplicationID(), registeredCommands)
	// slog.Error("error while registering commands", slog.Any("err", err))

	// if err = client.OpenGateway(context.TODO()); err != nil {
	// 	slog.Error("error while connecting to gateway", slog.Any("err", err))
	// }
}

// HandleCommand gère les interactions de commande
func HandleCommand(s bot.Client, i *events.ApplicationCommandInteractionCreate) {

	t := &events.IntegrationCreate{
		GenericIntegration: i,
	}

	if handler, exists := commandHandlers[t.SlashCommandInteractionData().CommandName()]; exists {
		handler(s, i)
	}
}
