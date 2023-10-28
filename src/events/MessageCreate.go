package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// var data = core.GetLanguage()

// func xpFetcher(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.GuildID == "" || message.Author.Bot || message.ChannelType != discordgo.ChannelTypeGuildText {
// 		return
// 	}

// 	baseData := database.GetUserXPLevelingData(message.GuildID, message.Author.ID)
// 	xpTurn := getGuildXPLevelingConfig(message.GuildID).Disable

// 	if xpTurn == "disable" {
// 		return
// 	}

// 	xp := baseData.XP
// 	level := baseData.Level
// 	randomNumber := rand.Intn(100) + 50

// 	database.SetUserXP(message.GuildID, message.Author.ID, "xp", xp+randomNumber)
// 	database.SetUserXP(message.GuildID, message.Author.ID, "xptotal", xp+randomNumber)

// 	if level*500 < xp {
// 		database.SetUserXP(message.GuildID, message.Author.ID, "level", level+1)
// 		database.SetUserXP(message.GuildID, message.Author.ID, "xp", xp-(level*500))

// 		newLevel := database.GetUserXPLevelingData(message.GuildID, message.Author.ID).Level

// 		if xpTurn == false {
// 			// Ou d'autres conditions pour afficher un message
// 			return
// 		}

// 		xpChan := getGuildXPLevelingConfig(message.GuildID).XPChannels

// 		if xpChan == "" {
// 			// Ou d'autres conditions pour afficher un message
// 			return
// 		}

// 		_, err := session.ChannelMessageSend(xpChan, data.EventXPLevelEarn.
// 			Replace("${message.author.id}", message.Author.ID).
// 			Replace("${newLevel}", newLevel))
// 		if err != nil {
// 			// Gérer les erreurs
// 		}
// 	}
// }

// func blockSpam(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.GuildID == "" || message.Channel == nil || message.Member == nil || message.ChannelType != discordgo.ChannelTypeGuildText || message.Author.Bot || message.Author.ID == session.State.User.ID {
// 		return
// 	}

// 	typeValue := getGuildAntipubConfig(message.GuildID)
// 	if typeValue == "off" || hasPermission(message.Member, discordgo.PermissionAdministrator) {
// 		return
// 	}

// 	member, _ := session.GuildMember(message.GuildID, message.Author.ID)

// 	if typeValue == "on" {
// 		logValue := getGuildPunishPubConfig(message.GuildID)
// 		logFetched := getTempPunishData(message.GuildID, message.Author.ID)

// 		if logValue.AmountMax == logFetched.Flags && logValue.State == "true" {
// 			switch logValue.PunishementType {
// 			case "ban":
// 				session.GuildBanCreateWithReason(message.GuildID, message.Author.ID, "Ban by PUNISHPUB", 0)
// 			case "kick":
// 				session.GuildMemberDeleteWithReason(message.GuildID, message.Author.ID, "Kick by PUNISHPUB")
// 			case "mute":
// 				muteRole := getMuteRole(message.GuildID)
// 				if muteRole != nil {
// 					session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, muteRole.ID)
// 					go func() {
// 						time.Sleep(40 * time.Second)
// 						if hasRole(member, muteRole.ID) {
// 							session.GuildMemberRoleRemove(message.GuildID, message.Author.ID, muteRole.ID)
// 						}
// 						setTempPunishData(message.GuildID, message.Author.ID, TempPunishData{})
// 					}()
// 				}
// 			}
// 		}

// 		blacklist := []string{"https://", "http://", "://", ".com", ".xyz", ".fr", "www.", ".gg", "g/", ".gg/", "youtube.be", "/?"}
// 		contentLower := strings.ToLower(message.Content)

// 		for _, word := range blacklist {
// 			if strings.Contains(contentLower, word) {
// 				flagsFetch := getTempPunishData(message.GuildID, message.Author.ID).Flags
// 				flagsFetch++
// 				setTempPunishData(message.GuildID, message.Author.ID, TempPunishData{Flags: flagsFetch})
// 				go func() {
// 					time.Sleep(40 * time.Second)
// 					session.ChannelMessageDelete(message.ChannelID, message.ID)
// 				}()
// 				break
// 			}
// 		}
// 	}
// }

// func rankRole(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	if message.GuildID == "" || message.Channel == nil || message.ChannelType != discordgo.ChannelTypeGuildText || message.Author.Bot ||
// 		message.Author.ID == session.State.User.ID || !hasPermission(message.Channel, discordgo.PermissionSendMessages) ||
// 		!hasPermission(message.Channel, discordgo.PermissionManageRoles) || message.Content != "<@"+session.State.User.ID+">" {
// 		return
// 	}

// 	roleID := getGuildRankRole(message.GuildID)
// 	role, err := session.GuildRole(message.GuildID, roleID)
// 	if err != nil {
// 		// Gérer les erreurs
// 		return
// 	}

// 	member, _ := session.GuildMember(message.GuildID, message.Author.ID)
// 	if hasRole(member, role.ID) {
// 		return
// 	}

// 	embed := &discordgo.MessageEmbed{
// 		Color: 0x4000ff,
// 		Title: "#" + role.ID,
// 		Author: &discordgo.MessageEmbedAuthor{
// 			Name:    data.EventRankRole.Replace("${message.author.username}", message.Author.Username),
// 			IconURL: message.Author.AvatarURL(""),
// 		},
// 		Description: "```" + message.Content + "```",
// 		Thumbnail: &discordgo.MessageEmbedThumbnail{
// 			URL: message.Guild.IconURL(),
// 		},
// 		Footer: &discordgo.MessageEmbedFooter{
// 			Text:    "iHorizon",
// 			IconURL: session.State.User.AvatarURL(""),
// 		},
// 		Timestamp: time.Now().Format(time.RFC3339),
// 	}

// 	session.GuildMemberRoleAdd(message.GuildID, message.Author.ID, role.ID)
// 	session.ChannelMessageSendEmbed(message.ChannelID, embed)
// }

// func createAllowList(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	baseData := getAllowList(message.GuildID)

// 	if baseData == nil {
// 		setAllowList(message.GuildID, &AllowListConfig{
// 			Enable: false,
// 			List: map[string]AllowListEntry{
// 				message.Guild.OwnerID: {Allowed: true},
// 			},
// 		})
// 	}
// }

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	reactToMessage(s, m)
}

func reactToMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID == "" || m.Author.Bot {
		return
	}

	recognizeItems := []string{
		"hey",
		"salut",
		"coucou",
		"bonjour",
		"salem",
		"wesh",
		"hello",
		"bienvenue",
	}

	contentLower := strings.ToLower(strings.Fields(m.Content)[0])
	for _, content := range recognizeItems {
		if strings.HasPrefix(contentLower, content) {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "👋")
			if err != nil {
				// Gérer les erreurs
			}
			return
		}
	}

}

// func suggestion(session *discordgo.Session, message *discordgo.MessageCreate) {
// 	baseData := getSuggestConfig(message.GuildID)

// 	if baseData == nil || baseData.Channel != message.ChannelID || baseData.Disable {
// 		return
// 	}

// 	suggestionContent := "```" + message.Content + "```"
// 	suggestCode := generateSuggestCode()

// 	suggestionEmbed := &discordgo.MessageEmbed{
// 		Color: 0x4000ff,
// 		Title: "#" + suggestCode,
// 		Author: &discordgo.MessageEmbedAuthor{
// 			Name:    data.EventSuggestionEmbedAuthor.Replace("${message.author.username}", message.Author.Username),
// 			IconURL: message.Author.AvatarURL(""),
// 		},
// 		Description: suggestionContent,
// 		Thumbnail: &discordgo.MessageEmbedThumbnail{
// 			URL: message.Guild.IconURL(),
// 		},
// 		Footer: &discordgo.MessageEmbedFooter{
// 			Text:    "iHorizon",
// 			IconURL: session.State.User.AvatarURL(""),
// 		},
// 		Timestamp: time.Now().Format(time.RFC3339),
// 	}

// 	session.ChannelMessageDelete(message.ChannelID, message.ID)

// 	args := strings.Fields(message.Content)
// 	if len(args) < 5 {
// 		return
// 	}

// 	msg, err := session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
// 		Content: "<@" + message.Author.ID + ">",
// 		Embed:   suggestionEmbed,
// 	})
// 	if err != nil {
// 		// Gérer les erreurs
// 		return
// 	}

// 	session.MessageReactionAdd(message.ChannelID, msg.ID, "✅")
// 	session.MessageReactionAdd(message.ChannelID, msg.ID, "❌")

// 	setSuggestionData(message.GuildID, suggestCode, &SuggestionData{
// 		Author: message.Author.ID,
// 		MsgID:  msg.ID,
// 	})
// }
