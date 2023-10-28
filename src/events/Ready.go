package events

import (
	commands "main/src"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	commands.RegisterCommands(s, s.State.User.ID, "999449972615413861")
}
