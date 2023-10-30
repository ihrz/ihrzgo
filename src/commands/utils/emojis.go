package utils

import (
	"fmt"
	commands "main/src"
	"main/src/core/lang"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "emojis",
		Description: "Generate a transgender-themed image",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "emojis",
				Description: "What the emoji then?",
				Required:    true,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		// Votre code ici
		var data = lang.GetLanguage(&i.GuildID)

		// if !hasAdministratorPermission(s, i.GuildID, i.User.ID) {
		// 	s.ChannelMessageSend(i.ChannelID, data["punishpub_not_admin"].(string))
		// 	return
		// }

		args := strings.Fields(options[0].StringValue())
		if len(args) < 2 {
			s.ChannelMessageSend(i.ChannelID, "Utilisation : !emojis <liste d'emojis>")
			return
		}

		cnt := 0
		nemj := ""
		for _, emoji := range args[1:] {
			match := regexp.MustCompile(`:(\w+):(\d+)>`).FindStringSubmatch(emoji)
			if len(match) == 3 {
				isAnimated := strings.HasPrefix(emoji, "<a:")
				emojiName := match[1]
				emojiID := match[2]

				err := createEmoji(s, i.GuildID, emojiName, emojiID, isAnimated)
				if err == nil {
					s.ChannelMessageSend(i.ChannelID, fmt.Sprintf(data["emoji_send_new_emoji"].(string), emojiName, emoji))
					cnt++
					nemj += fmt.Sprintf("<%s:%s:%s>", getEmojiPrefix(isAnimated), emojiName, emojiID)
				} else {
					s.ChannelMessageSend(i.ChannelID, fmt.Sprintf(data["emoji_send_err_emoji"].(string), emojiName))
				}
			}
		}

		embed := &discordgo.MessageEmbed{
			Color:       0xbea9de,
			Footer:      &discordgo.MessageEmbedFooter{Text: "iHorizon", IconURL: s.State.User.AvatarURL("512")},
			Timestamp:   time.Now().Format("2006-01-02T15:04:05.000Z"),
			Description: fmt.Sprintf(data["emoji_embed_desc_work"].(string), cnt, i.GuildID, nemj),
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}

func hasAdministratorPermission(s *discordgo.Session, guildID string, userID string) bool {
	member, err := s.GuildMember(guildID, userID)
	if err != nil {
		return false
	}
	return (member.Permissions & discordgo.PermissionAdministrator) != 0
}

func createEmoji(s *discordgo.Session, guildID string, emojiName string, emojiID string, isAnimated bool) error {
	emojiURL := fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.%s", emojiID, map[bool]string{true: "gif", false: "png"}[isAnimated])

	params := &discordgo.EmojiParams{
		Name:  emojiName,
		Image: emojiURL,
	}

	_, err := s.GuildEmojiCreate(guildID, params)
	return err
}

func getEmojiPrefix(isAnimated bool) string {
	if isAnimated {
		return "a"
	}
	return ""
}
