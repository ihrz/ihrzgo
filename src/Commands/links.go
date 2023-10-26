package commands

import (
	commands "main/src"
	"main/src/core"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// Enregistrez la commande
	command := &discordgo.ApplicationCommand{
		Name:        "links",
		Description: "Get all links about iHorizon!",
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		core.GetLanguage()
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Powng!",
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}
