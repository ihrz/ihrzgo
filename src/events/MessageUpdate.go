package events

import (
	"main/src/core/lang"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	serverLogsForMessageUpdate(s, m)
}

func serverLogsForMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {

	if m.Author == nil || m.Author.Bot || len(m.Content) < 1 || m.BeforeUpdate == nil {
		return
	}

	var data = lang.GetLanguage(&m.GuildID)

	// someInfo := db.GetServerLogsMessage("msg", m.GuildID)

	// if someInfo == "" || m.Content == "" || m.BeforeUpdate.Content == m.Content {
	// 	return
	// }

	msgChannel, err := s.State.Channel("1139545233776975975")
	if err != nil || msgChannel == nil {
		return
	}

	icon := m.Author.AvatarURL("")

	logsEmbed := &discordgo.MessageEmbed{
		Color: 0x000000,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: icon,
		},
		Description: strings.Replace(data["event_srvLogs_messageUpdate_description"].(string), "${oldMessage.channelId}", m.ChannelID, -1),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  data["event_srvLogs_messageUpdate_footer_1"].(string),
				Value: " " + m.BeforeUpdate.Content,
			},
			{
				Name:  data["event_srvLogs_messageUpdate_footer_2"].(string),
				Value: " " + m.Content,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = s.ChannelMessageSendComplex(msgChannel.ID, &discordgo.MessageSend{Embed: logsEmbed})
	if err != nil {
		return
	}

}

// func getServerLogsMessage(guildID string) string {
// 	// Mettez ici la logique pour récupérer l'ID du canal de logs
// 	return ""
// }
