package events

import (
	"fmt"
	"main/src/core/db"
	"main/src/core/lang"

	"github.com/bwmarrin/discordgo"
)

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	println("a Message ahs been deleted")
	lang := lang.GetLanguage(&m.GuildID)
	_ = lang["event_srvLogs_messageDelete_description"].(string)
	if m.Author != nil {
		_ = m.Author.AvatarURL("512")
	}

	_ = m.Timestamp
	id_channel_logs := db.GetLogsChannel(&m.ChannelID)
	if len(m.Attachments) > 1 {
		return
	}
	if len(m.Content) < 1 {
		//message vide
		return
	}
	_, err := s.ChannelMessageSend(id_channel_logs, m.Content)
	if err != nil {
		fmt.Println(err.Error())
	}

}
