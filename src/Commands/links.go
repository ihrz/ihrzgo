package commands

import (
	commands "main/src"
	"main/src/core/lang"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// Enregistrez la commande
	command := &discordgo.ApplicationCommand{
		Name:        "links",
		Description: "Get all links about iHorizon!",
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		var lang = lang.GetLanguage(&i.GuildID)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: lang["links_message"].(string),
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}
