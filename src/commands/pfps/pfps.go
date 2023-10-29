package pfps

import (
	commands "main/src"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "pfps",
		Description: "Manage PFPS settings",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "disable",
				Description: "Enable or Disable the module!",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "action",
						Description: "What do you want to do ?",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "Power On",
								Value: "on",
							},
							{
								Name:  "Power Off",
								Value: "off",
							},
						},
					},
				},
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// var data = lang.GetLanguage(&i.GuildID)

		options := i.ApplicationCommandData().Options

		if options[0].Name == "disable" {
			// Votre logique pour désactiver la fonctionnalité PFPS dans le canal
			// Assurez-vous de gérer correctement les erreurs et la logique métier ici.
			// Vous pouvez utiliser i.GuildID, i.Member, etc. pour accéder aux informations de l'interaction.

			// Par exemple, pour répondre à l'interaction :
			response := "PFPS feature has been disabled in this channel."
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: response,
				},
			})
		}
	}

	commands.AddCommand(command, commandHandler)
}
