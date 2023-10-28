package events

import (
	"fmt"
	"main/src/core/config"

	"github.com/bwmarrin/discordgo"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	// Gestionnaire d'événement pour la suppression d'un serveur
	if g.Guild.Name == "" || g.Guild.ID == "" {
		return
	}

	var i = "None"
	if g.Guild.VanityURLCode != "" {
		i = "discord.gg/" + g.Guild.VanityURLCode
	}

	embed := &discordgo.MessageEmbed{
		Color:       0xff0505,
		Timestamp:   g.JoinedAt.Format(""),
		Description: fmt.Sprintf("**A guild have deleted iHorizon !**"),
		Fields: []*discordgo.MessageEmbedField{
			{Name: "🏷️・Server Name", Value: "`" + g.Guild.Name + "`", Inline: true},
			{Name: "🆔・Server ID", Value: "`" + g.Guild.ID + "`", Inline: true},
			{Name: "🌐・Server Region", Value: "`" + g.Guild.PreferredLocale + "`", Inline: true},
			{Name: "👤・Member Count", Value: "`" + fmt.Sprint(g.Guild.MemberCount) + " members`", Inline: true},
			{Name: "🪝・Vanity URL", Value: "`" + i + "`", Inline: true},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: g.Guild.IconURL(""),
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "iHorizon",
			IconURL: "URL_DE_VOTRE_ICONE",
		},
	}

	channel, err := s.State.Channel(config.Main.Core.GuildLogsChannelID)

	if err != nil {
		fmt.Println("Erreur lors de la récupération du canal:", err)
		return
	}

	_, err = s.ChannelMessageSendEmbed(channel.ID, embed)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du message:", err)
		return
	}
}
