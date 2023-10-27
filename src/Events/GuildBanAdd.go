package events

import (
	"main/src/core/lang"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func GuildBanAdd(s *discordgo.Session, b *discordgo.GuildBanAdd) {
	serverLogs(s, b)
}

func serverLogs(session *discordgo.Session, ban *discordgo.GuildBanAdd) {
	// p, err := session.UserChannelPermissions(session.State.User.ID, ban.GuildID)

	// if p&discordgo.PermissionViewAuditLogs != discordgo.PermissionViewAuditLogs {
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

	description := strings.Replace(data["event_srvLogs_banAdd_description"].(string), "${firstEntry.executor.id}", firstEntry.ID, -1)
	description = strings.Replace(description, "${firstEntry.target.id}", ban.User.ID, -1)

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
