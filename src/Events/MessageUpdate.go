package events

import (
	"github.com/bwmarrin/discordgo"
)

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	serverLogsForMessageUpdate(s, m)
}

func serverLogsForMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	println(m, s)
	println(m.BeforeUpdate.Content)
}

// func getServerLogsMessage(guildID string) string {
// 	// Mettez ici la logique pour récupérer l'ID du canal de logs
// 	return ""
// }

// message, err := s.State.Message(m.ChannelID, m.ID)
// if err != nil {
// 	fmt.Println("Error when atempt to cache this message :", err)
// 	return
// }

// fmt.Printf("%+v\n", message)

// if m.Author == nil || m.Author.Bot || m.Content == "" {
// 	return
// }

// var data = lang.GetLanguage()
// someInfo := getServerLogsMessage(oldMessage.GuildID)

// if someInfo == "" || oldMessage.Content == "" || newMessage.Content == "" || oldMessage.Content == newMessage.Content {
// 	return
// }

// msgChannel, err := session.State.Channel(someInfo)
// if err != nil || msgChannel == nil {
// 	return
// }

// icon := newMessage.Author.AvatarURL("")

// logsEmbed := &discordgo.MessageEmbed{
// 	Color: 0x000000,
// 	Author: &discordgo.MessageEmbedAuthor{
// 		Name:    newMessage.Author.Username,
// 		IconURL: icon,
// 	},
// 	Description: fmt.Sprintf(data.EventSrvLogsMessageUpdateDescription, oldMessage.ChannelID),
// 	Fields: []*discordgo.MessageEmbedField{
// 		{
// 			Name:  data.EventSrvLogsMessageUpdateFooter1,
// 			Value: " " + oldMessage.Content,
// 		},
// 		{
// 			Name:  data.EventSrvLogsMessageUpdateFooter2,
// 			Value: " " + newMessage.Content,
// 		},
// 	},
// 	Timestamp: time.Now().Format(time.RFC3339),
// }

// _, err = session.ChannelMessageSendComplex(msgChannel.ID, &discordgo.MessageSend{Embed: logsEmbed})
// if err != nil {
// 	// Gérer les erreurs
// }
