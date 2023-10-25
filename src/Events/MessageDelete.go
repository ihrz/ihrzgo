package events

import "github.com/bwmarrin/discordgo"

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	println("message supprimé: {m.Content}")
}
