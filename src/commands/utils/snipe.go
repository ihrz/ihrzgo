package utils

import (
	"fmt"
	commands "main/src"
	"main/src/core/db"
	"main/src/core/lang"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "snipe",
		Description: "Get the last message deleted in this channel!",
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := lang.GetLanguage(&i.GuildID)

		based := db.GetLastMessageDeleteHere(&i.GuildID, &i.ChannelID)

		if based.Snipe == "" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: data["snipe_no_previous_message_deleted"].(string),
				},
			})
			return
		}

		embed := &discordgo.MessageEmbed{
			Color: 0x474749,
			// Author: &discordgo.MessageEmbedAuthor{
			// 	Name:    based.SnipeUserInfoTag,
			// 	IconURL: based.SnipeUserInfoPp,
			// },
			Description: fmt.Sprintf("`%s`", based.Snipe),
			// Timestamp:   based.SnipeTimestamp,
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})

		if err != nil {
			panic(err)
		}
	}

	commands.AddCommand(command, commandHandler)
}
