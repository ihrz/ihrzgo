package fun

import (
	commands "main/src"
	"main/src/core/lang"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "love",
		Description: "Calculate the love percentage between two users",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user1",
				Description: "The first user",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user2",
				Description: "The second user",
				Required:    false,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var lang = lang.GetLanguage(&i.GuildID)
		options := i.ApplicationCommandData().Options

		user1ID := options[0].UserValue(s).ID
		user2ID := options[1].UserValue(s).ID

		var user1, user2 *discordgo.User

		if user1ID != "" {
			user1, _ = s.User(user1ID)
		} else {
			user1 = i.Member.User
		}

		if user2ID != "" {
			user2, _ = s.User(user2ID)
		} else {
			// Remplacer par la logique pour obtenir un utilisateur au hasard dans la guilde
		}

		// profileImageSize := 512
		// canvasWidth := profileImageSize * 3
		// canvasHeight := profileImageSize

		// dc := gg.NewContext(canvasWidth, canvasHeight)
		// heartEmojiPath := "./src/assets/heart.png"
		// profileImage1URL := user1.AvatarURL("512")
		// profileImage2URL := user2.AvatarURL("512")

		// Charger les images

		// profileImage1 := loadImage(profileImage1URL)
		// profileImage2 := loadImage(profileImage2URL)
		// heartEmoji := loadImage(heartEmojiPath)

		// dc.DrawImage(profileImage1, 0, 0)
		// heartX := float64(profileImageSize)
		// heartY := float64(profileImageSize/2 - heartEmoji.Height()/2)
		// dc.DrawImage(heartEmoji, int(heartX), int(heartY))
		// dc.DrawImage(profileImage2, int(profileImageSize*1+float64(heartEmoji.Width())), 0)

		// Convertir le dessin en une image

		// img := dc.Image()
		// img = resize.Resize(512, 0, img, resize.Lanczos3)
		// buffer := new(bytes.Buffer)
		// png.Encode(buffer, img)

		// always100 := config.Main.More.Always100
		// user1IDStr := user1.ID
		// user2IDStr := user2.ID
		lovePercentage := 0

		// always100 := config.Main.More.Always100
		// user1IDStr := user1.ID
		// user2IDStr := user2.ID

		// if contains(always100, user1IDStr+"x"+user2IDStr) || contains(always100, user2IDStr+"x"+user1IDStr) {
		// 	lovePercentage = 100
		// } else {
		// }

		lovePercentage = rand.Intn(101)

		desc := strings.Replace(lang["love_embed_description"].(string), "${user1.username}", user1.Username, -1)
		desc = strings.Replace(desc, "${user2.username}", user2.Username, -1)
		desc = strings.Replace(desc, "${randomNumber}", strconv.Itoa(lovePercentage), -1)

		embed := &discordgo.MessageEmbed{
			Color:       0xFFC0CB,
			Title:       "💕",
			Description: desc,
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "iHorizon",
				IconURL: s.State.User.AvatarURL(""),
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		// Répondre à l'interaction avec l'image et l'embed
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
			// Files: []*discordgo.File{
			// 	{
			// 		Name:   "love.png",
			// 		Reader: buffer,
			// 	},
			// },
		})
	}

	commands.AddCommand(command, commandHandler)
}

// Helper function to check if a string is in a slice
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// // Helper function to load an image from a URL
// func loadImage(url string) *gg.Image {
// 	image, err := gg.LoadImage(url)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return image
// }
