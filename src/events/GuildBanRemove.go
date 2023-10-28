package events

import (
	"main/src/core/lang"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GuildBanRemove(session *discordgo.Session, ban *discordgo.GuildBanRemove) {
	serverLogsForBanRemove(session, ban)
}

func serverLogsForBanRemove(session *discordgo.Session, ban *discordgo.GuildBanRemove) {
	// if session.State.User == nil || !hasPermission(session.State.User.ID, ban.GuildID, discordgo.PermissionViewAuditLogs) {
	// 	return
	// }

	fetchedLogs, err := session.GuildAuditLog(ban.GuildID, "", "", int(discordgo.AuditLogActionMemberBanAdd), 1)
	if err != nil {
		// Gérer les erreurs
		return
	}

	if len(fetchedLogs.Users) == 0 {
		return
	}

	firstEntry := fetchedLogs.Users[0]
	// someInfo := getServerLogsModeration(ban.GuildID)

	// if someInfo == "" {
	// 	return
	// }
	var data = lang.GetLanguage()

	msgChannel, err := session.State.Channel("1139545236268384356")
	if err != nil {
		// Gérer les erreurs
		return
	}

	description := strings.Replace(data["event_srvLogs_banRemove_description"].(string), "${firstEntry.executor.id}", firstEntry.ID, -1)
	description = strings.Replace(description, "${firstEntry.target.username}", firstEntry.Username, -1)

	logsEmbed := &discordgo.MessageEmbed{
		Color:       0x000000,
		Description: description,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	_, err = session.ChannelMessageSendComplex(msgChannel.ID, &discordgo.MessageSend{Embed: logsEmbed})
	if err != nil {
		// Gérer les erreurs
	}
}
