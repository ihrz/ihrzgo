package module

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(dg *discordgo.Session, guildID, channelID string) {
	// guild, err := dg.Guild(guildID)
	// if err != nil {
	// 	fmt.Println("error fetching guild,", err)
	// 	return
	// }

	channel, err := dg.State.Channel(channelID)
	if err != nil {
		fmt.Println("error fetching channel,", err)
		return
	}

	if channel.Type != discordgo.ChannelTypeGuildText {
		fmt.Println("channel is not a text channel")
		return
	}

	members, err := dg.GuildMembers(guildID, "", 1000)
	if err != nil {
		fmt.Println("error fetching guild members,", err)
		return
	}

	// Shuffle members to prevent the same user from being selected consecutively
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(members), func(i, j int) {
		members[i], members[j] = members[j], members[i]
	})

	var user *discordgo.Member
	// var lastUser *discordgo.Member

	usr := ""
	for _, m := range members {
		if m.User.Bot || m.User.ID == usr {
			continue
		}
		usr = m.User.ID
		user = m
		break
	}

	if user == nil {
		fmt.Println("no suitable user found")
		return
	}

	actRow := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Style: discordgo.LinkButton,
					URL:   user.User.AvatarURL(""),
					Label: "Download User Avatar",
				},
			},
		},
	}

	ebds := []*discordgo.MessageEmbed{}

	if user.User.Avatar != "" {
		ebds = append(ebds, &discordgo.MessageEmbed{
			Color: 0xa2add0,
			Title: fmt.Sprintf("%s's **User** avatar", user.User.Username),
			Image: &discordgo.MessageEmbedImage{
				URL: user.User.AvatarURL(""),
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Footer: &discordgo.MessageEmbedFooter{
				Text: "iHorizon",
			},
		})
	}

	_, err = dg.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
		Embeds:     ebds,
		Components: actRow,
	})
	if err != nil {
		fmt.Println("error sending message,", err)
		return
	}
}
