package utils

import (
	commands "main/src"
	"main/src/core/lang"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "renew",
		Description: "Re-created a channels (cloning permission and all configurations). nuke equivalent",
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var data = lang.GetLanguage(&i.GuildID)

		if !HasAdministratorPermission(int(i.Member.Permissions)) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: data["renew_not_administrator"].(string),
				},
			})
			return
		}

		channel, err := s.Channel(i.ChannelID)
		if err != nil {
			// Gérer l'erreur
			return
		}

		channel, err = s.ChannelDelete(channel.ID)
		if err != nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: data["renew_dont_have_permission"].(string),
				},
			})
			return
		}

		here, err := s.GuildChannelCreate(channel.GuildID, channel.Name, discordgo.ChannelTypeGuildText)
		if err != nil {
			// Gérer l'erreur
			return
		}

		// Appliquer les anciennes autorisations au nouveau canal
		// for _, permission := range oldPermissions {
		// 	// _, err := s.ChannelPermissionSet(here.ID, permission.ID, permission.Type, permission.Allow, permission.Deny)
		// 	if err != nil {
		// 		// Gérer l'erreur
		// 		return
		// 	}
		// }
		successMessage := strings.Replace(data["renew_channel_send_success"].(string), "${interaction.user}", "<@"+i.Member.User.ID+">", -1)
		s.ChannelMessageSend(here.ID, successMessage)
	}

	commands.AddCommand(command, commandHandler)
}

func HasAdministratorPermission(permissions int) bool {
	return (permissions & discordgo.PermissionAdministrator) != 0
}
