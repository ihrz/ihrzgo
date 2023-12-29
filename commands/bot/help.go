package bot

import (
	"log/slog"
	handlers "main/handlers"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func init() {
	// Enregistrez la commande
	command := &discord.SlashCommandCreate{
		Name:        "help",
		Description: "Get list of all command !",
	}

	commandHandler := func(s bot.Client, i *events.ApplicationCommandInteractionCreate) {

		embed := discord.Embed{
			Title:       "Help Panel",
			Description: "coucou",
			Fields: []discord.EmbedField{
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
			Footer: &discord.EmbedFooter{
				Text: i.Client().ID().String(),
				// IconURL: s.State.User.AvatarURL(""),
			},
			Color: 15859878,
		}

		err := i.CreateMessage(discord.NewMessageCreateBuilder().SetEmbeds(embed).Build())

		if err != nil {
			slog.Error("error on sending response", slog.Any("err", err))
		}
		// s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// 	Type: discordgo.InteractionResponseChannelMessageWithSource,
		// 	Data: &discordgo.InteractionResponseData{
		// 		Embeds: []*discordgo.MessageEmbed{embed},
		// 	},
		// })
	}

	handlers.AddCommand(command, commandHandler)
}

func Build() {
	panic("unimplemented")
}
