package events

import (
	"main/src/core/lang"

	"github.com/bwmarrin/discordgo"
)

func ChannelCreate(s *discordgo.Session, m *discordgo.ChannelCreate) {
	ihrzLogs(s, m)
}

func ihrzLogs(session *discordgo.Session, channel *discordgo.ChannelCreate) {
	if channel.Name != "ihorizon-logs" {
		return
	}

	lang := lang.GetLanguage()

	setupEmbed := &discordgo.MessageEmbed{
		Color:       0x1e1d22,
		Title:       lang["event_channel_create_message_embed_title"].(string),
		Description: lang["event_channel_create_message_embed_description"].(string),
	}

	_, err := session.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{Embed: setupEmbed})
	if err != nil {
		panic(err)
	}
}
