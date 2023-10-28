package fun

import (
	"bytes"
	commands "main/src"

	"github.com/bwmarrin/discordgo"
	resty "github.com/go-resty/resty/v2"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "transgender",
		Description: "Generate a transgender-themed image",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "User to use as an avatar (optional)",
				Required:    false,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		userOption := options[0].UserValue(s)
		user := i.Member.User
		if userOption != nil {
			user = userOption
		}

		link := generateTransgenderImageURL(user.AvatarURL("png"))

		imgs, err := getTransgenderImage(link)
		if err != nil {
			// Handle the error
			return
		}

		embed := &discordgo.MessageEmbed{
			Color: 0x000000,
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://all-humans-have-right-elektra.png",
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "iHorizon x ElektraBots",
				IconURL: s.State.User.AvatarURL(""),
			},
			// Timestamp: i.CreatedAt,
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
				Files:  []*discordgo.File{imgs},
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}

func generateTransgenderImageURL(avatarURL string) string {
	return "https://some-random-api.com/canvas/misc/transgender?avatar=" + avatarURL
}

func getTransgenderImage(url string) (*discordgo.File, error) {
	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	imgs := &discordgo.File{
		Name:   "all-humans-have-right-elektra.png",
		Reader: bytes.NewReader(resp.Body()),
	}

	return imgs, nil
}
