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

		var lang = core.GetLanguage()

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: lang["links_msg"].(string),
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}
