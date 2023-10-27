package events

import (
	"main/src/core/config"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	messageToServer(s, g)
}

func messageToServer(session *discordgo.Session, guild *discordgo.GuildCreate) {
	welcomeMessages := []string{
		"Welcome to our server! 🎉", "Greetings, fellow Discordians! 👋",
		"iHorizon has joined the chat! 💬", "It's a bird, it's a plane, no, it's iHorizon! 🦸‍♂",
		"Let's give a warm welcome to iHorizon! 🔥",
	}

	welcomeMessage := welcomeMessages[rand.Intn(len(welcomeMessages))]

	embed := &discordgo.MessageEmbed{
		Color:     0x00FF00,
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     welcomeMessage,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "iHorizon",
			IconURL: session.State.User.AvatarURL(""),
		},
		Description: `Hi there! I'm excited to join your server and be a part of your community. 

My name is iHorizon, and I'm here to help you with all your needs. Feel free to use my commands and explore all the features I have to offer.

If you have any questions or run into any issues, don't hesitate to reach out to me.
I'm here to make your experience on this server the best it can be. 

Thanks for choosing me, and let's have some fun together!`,
	}

	if channel, err := session.State.Channel(guild.Guild.SystemChannelID); err == nil {
		session.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{Embed: embed})
	}
}

func ownerLogs(session *discordgo.Session, guild *discordgo.Guild) {
	var i string
	if guild.VanityURLCode != "" {
		i = "discord.gg/" + guild.VanityURLCode
	}

	var channel *discordgo.Channel

	if guild.SystemChannelID != "" {
		channel, _ = session.State.Channel(guild.SystemChannelID)
	}

	if channel == nil {
		// Si le canal système n'est pas défini, choisissez un canal au hasard
		channels := guild.Channels
		if len(channels) > 0 {
			rand.Seed(time.Now().UnixNano())
			channel = channels[rand.Intn(len(channels))]
		}
	}

	inviteLink, err := createInvite(session, channel)
	if err != nil {
		inviteLink = "None"
	}

	embed := &discordgo.MessageEmbed{
		Color:       0x00FF00,
		Timestamp:   guild.JoinedAt.Format(time.RFC3339),
		Description: "**A new guild have added iHorizon !**",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "🏷️・Server Name",
				Value:  "`" + guild.Name + "`",
				Inline: true,
			},
			{
				Name:   "🆔・Server ID",
				Value:  "`" + guild.ID + "`",
				Inline: true,
			},
			{
				Name:   "🌐・Server Region",
				Value:  "`" + guild.PreferredLocale + "`",
				Inline: true,
			},
			{
				Name:   "👤・Member Count",
				Value:  "`" + strconv.Itoa(guild.MemberCount) + " members`",
				Inline: true,
			},
			{
				Name:   "🔗・Invite Link",
				Value:  "`" + inviteLink + "`",
				Inline: true,
			},
			{
				Name:   "🪝・Vanity URL",
				Value:  "`" + i + "`",
				Inline: true,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: guild.IconURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "iHorizon",
			IconURL: session.State.User.AvatarURL(""),
		},
	}

	_, err = session.ChannelMessageSendComplex(config.Main.Core.GuildLogsChannelID, &discordgo.MessageSend{Embed: embed})
	if err != nil {
		// Gérer les erreurs
	}
}

func createInvite(session *discordgo.Session, channel *discordgo.Channel) (string, error) {
	invite, err := session.ChannelInviteCreate(channel.ID, discordgo.Invite{
		MaxUses: 1,
		Unique:  true,
	})
	if err != nil {
		return "None", err
	}
	return "discord.gg/" + invite.Code, nil
}
