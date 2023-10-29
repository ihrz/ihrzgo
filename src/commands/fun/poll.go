package fun

import (
	"fmt"
	commands "main/src"
	"main/src/core/lang"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "poll",
		Description: "Create a poll by reacting with ✅ and ❌",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Message for the poll",
				Required:    true,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		pollMessage := options[0].StringValue()

		var data = lang.GetLanguage(&i.GuildID)

		if !hasAdministratorPermissions(i.Member.Permissions) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: data["poll_not_admin"].(string),
				},
			})
			return
		}

		title := strings.Replace(data["poll_embed_title"].(string), "${interaction.user.username}", i.Member.User.Username, -1)

		pollEmbed := &discordgo.MessageEmbed{
			Title:       title,
			Color:       0xddd98b,
			Description: pollMessage,
			Fields: []*discordgo.MessageEmbedField{
				{Name: data["poll_embed_fields_reaction"].(string), Value: data["poll_embed_fields_choice"].(string)},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "https://cdn.discordapp.com/attachments/610152915063013376/610947097969164310/loading-animation.gif",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		// Répondre à l'interaction avec l'incorporation
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{pollEmbed},
			},
		})
		if err != nil {
			// Gérez l'erreur
			return
		}

		// Ajouter les réactions "✅" et "❌" au message
		msg, err := s.InteractionResponse(i.Interaction)

		err = s.MessageReactionAdd(i.ChannelID, msg.ID, "✅")
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}

		err = s.MessageReactionAdd(i.ChannelID, msg.ID, "❌")
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	}

	commands.AddCommand(command, commandHandler)
}

func hasAdministratorPermissions(permissions int64) bool {
	return (permissions & discordgo.PermissionAdministrator) != 0
}
