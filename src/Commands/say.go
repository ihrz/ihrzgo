// ping.go
package commands

import (
	commands "main/src"

	"github.com/bwmarrin/discordgo"
)

func init() {
	// Enregistrez la commande
	command := &discordgo.ApplicationCommand{
		Name:        "say",
		Description: "say command!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "content",
				Description: "The message to sent",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		content, ok := options[0].Value.(string)

		if !ok {
			// Gérez l'erreur en conséquence si la conversion échoue
			return
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}
