package fun

import (
	commands "main/src"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var alpha = []string{
	" ", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	"T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3",
	"4", "5", "6", "7", "8", "9", "0",
}

var morse = []string{
	"/", ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
	".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...",
	"-", "..-", "...-", ".--", "-..-", "-.--", "--..", ".----", "..---", "...--",
	"....-", ".....", "-....", "--...", "---..", "----.", "-----",
}

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "morse",
		Description: "Convert text to/from Morse code",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "input",
				Description: "Text to convert",
				Required:    true,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		input := options[0].StringValue()
		text := strings.ToUpper(input)

		for strings.Contains(text, "Ä") || strings.Contains(text, "Ö") || strings.Contains(text, "Ü") {
			text = strings.Replace(text, "Ä", "AE", -1)
			text = strings.Replace(text, "Ö", "OE", -1)
			text = strings.Replace(text, "Ü", "UE", -1)
		}

		var convertedText string

		if strings.HasPrefix(text, ".") || strings.HasPrefix(text, "-") {
			text = strings.Replace(text, " ", "", -1)
			letters := strings.Split(text, " ")
			for _, letter := range letters {
				index := indexOf(morse, letter)
				if index >= 0 && index < len(alpha) {
					convertedText += alpha[index]
				}
			}
		} else {
			letters := strings.Split(text, "")
			for _, letter := range letters {
				index := indexOf(alpha, letter)
				if index >= 0 && index < len(morse) {
					convertedText += morse[index] + " "
				}
			}
		}

		// Enlever l'espace final
		convertedText = strings.TrimSpace(convertedText)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "```" + convertedText + "```",
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}

func indexOf(arr []string, value string) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
