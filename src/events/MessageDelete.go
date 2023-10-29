package events

import (
	"main/src/core/db"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type SnipeChannelStruc struct {
	SnipeUserInfoTag string
	SnipeUserInfoPp  string
	Snipe            string
	SnipeTimestamp   string
}

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	// lang := lang.GetLanguage(&m.GuildID)
	// _ = lang["event_srvLogs_messageDelete_description"].(string)
	// if m.Author != nil {
	// 	_ = m.Author.AvatarURL("512")
	// }

	// _ = m.Timestamp
	// id_channel_logs := db.GetLogsChannel(&m.ChannelID)
	// if len(m.Attachments) > 1 {
	// 	return
	// }
	// if len(m.Content) < 1 {
	// 	//message vide
	// 	return
	// }
	// _, err := s.ChannelMessageSend(id_channel_logs, m.Content)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	snipeModule(s, m)
}

func snipeModule(s *discordgo.Session, m *discordgo.MessageDelete) {
	var db = db.GetDatabase()

	str := strings.Replace("${m.GuildID}_SNIPE_${m.ChannelID}", "${m.GuildID}", m.GuildID, -1)
	str = strings.Replace(str, "${m.ChannelID}", m.ChannelID, -1)

	db.AutoMigrate(&SnipeChannelStruc{})

	db.Create(&SnipeChannelStruc{SnipeUserInfoTag: m.Author.Username, Snipe: m.BeforeDelete.Content})
}
