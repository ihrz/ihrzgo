package fun

import (
	commands "main/src"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := &discordgo.ApplicationCommand{
		Name:        "subcommands",
		Description: "Subcommands and command groups example",
		Options: []*discordgo.ApplicationCommandOption{

			{
				Name:        "subcommand-group",
				Description: "Subcommands group",
				Options: []*discordgo.ApplicationCommandOption{
					// Also, subcommand groups aren't capable of
					// containing options, by the name of them, you can see
					// they can only contain subcommands
					{
						Name:        "nested-subcommand",
						Description: "Nested subcommand",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
				},
				Type: discordgo.ApplicationCommandOptionSubCommandGroup,
			},
			// Also, you can create both subcommand groups and subcommands
			// in the command at the same time. But, there's some limits to
			// nesting, count of subcommands (top level and nested) and options.
			// Read the intro of slash-commands docs on Discord dev portal
			// to get more information
			{
				Name:        "subcommand",
				Description: "Top-level subcommand",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	}

	commandHandler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	}

	commands.AddCommand(command, commandHandler)
}
