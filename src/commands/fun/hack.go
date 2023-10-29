package fun

import (
	commands "main/src"
	"main/src/core/lang"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var ip = []string{
	"1", "100", "168", "254", "345", "128", "256", "255", "0", "144",
	"38", "67", "97", "32", "64", "192", "10", "172", "12", "200", "87",
	"150", "42", "99", "76", "211", "172", "18", "86", "55", "220", "7",
}

var hackerNames = []string{
	"cyberpunk", "zeroday", "blackhat", "hackmaster", "shadowbyte", "crypt0",
	"phishr", "darknet", "rootaccess", "sploit3r", "hack3rman", "v1rus",
	"bytebandit", "malware", "scriptkiddie",
}

var hackerDomains = []string{
	"hackmail.com", "darkweb.net", "blackhat.org", "zerodaymail.com",
	"phishmail.net", "cryptomail.org", "sploitmail.com", "hackergang.com",
	"rootmail.org", "v1rusmail.com",
}

var hackerPasswords = []string{
	"5up3rP@$$w0rd", "H4x0r!z3d",
	"N0s3cur1ty", "3vilG3nius", "0bscureC0de", "Hacker123!", "P@$$phr4s3",
	"D3c3pt10n", "0v3rwr1t3", "V1rtu4lInf1ltr4t0r", "R3v3rse3ng1n33r",
	"C0mpl3xM4tr1x", "D1g1t4lS3cr3t", "Myst3ryH4ck", "Ph4nt0mC0ntrol",
}

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "hack",
		Description: "Generate a random hacking scenario",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The victim to target",
				Required:    true,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var data = lang.GetLanguage(&i.GuildID)

		options := i.ApplicationCommandData().Options

		victim := options[0].UserValue(s)
		if victim == nil {
			return
		}

		generatedIp := generateRandomIp()
		generatedUsername := generateRandomUsername()
		generatedEmail := generateRandomEmail(generatedUsername)
		generatedPassword := generateRandomPassword()

		embed := &discordgo.MessageEmbed{
			Color:       0x800000,
			Description: data["hack_embed_description"].(string),
			Fields: []*discordgo.MessageEmbedField{
				{Name: data["hack_embed_fields_ip"].(string), Value: "`" + generatedIp + "`"},
				{Name: data["hack_embed_fields_email"].(string), Value: "`" + generatedEmail + "`"},
				{Name: data["hack_embed_fields_password"].(string), Value: "`" + generatedPassword + "`"},
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		// Replace placeholders in the description with actual values
		embed.Description = strings.Replace(embed.Description, "${victim.id}", victim.ID, -1)
		embed.Description = strings.Replace(embed.Description, "${interaction.user.id}", i.Member.User.ID, -1)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}

	commands.AddCommand(command, commandHandler)
}

func generateRandomIp() string {
	return strings.Join([]string{
		ip[rand.Intn(len(ip))+1],
		ip[rand.Intn(len(ip))+1],
		ip[rand.Intn(len(ip))+1],
		ip[rand.Intn(len(ip))+1],
	}, ".")
}

func generateRandomUsername() string {
	return hackerNames[rand.Intn(len(hackerNames))] + generateRandomNumber()
}

func generateRandomEmail(username string) string {
	return username + "@" + hackerDomains[rand.Intn(len(hackerDomains))]
}

func generateRandomPassword() string {
	return hackerPasswords[rand.Intn(len(hackerPasswords))]
}

func generateRandomNumber() string {
	possible := "0123456789"
	var builder strings.Builder
	for i := 0; i < 8; i++ {
		builder.WriteByte(possible[rand.Intn(len(possible))])
	}
	return builder.String()
}
