// ping.go
package commands

import (
	commands "main/src"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// Enregistrez la commande
	command := &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Get list of all command !",
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		embed := &discordgo.MessageEmbed{
			Title:       "Help Panel",
			Description: "coucou",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Users",
					Value: "test",
				},
				{
					Name:  "Channels",
					Value: "tes2",
				},
				{
					Name:  "Roles",
					Value: "tes3",
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    s.State.User.Username,
				IconURL: s.State.User.AvatarURL(""),
			},
			Color: 15859878,
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}
