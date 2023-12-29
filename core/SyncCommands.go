package core

import (
	bot_2 "main/commands/bot"
	commands "main/handlers"

	"github.com/disgoorg/disgo/bot"
)

func SyncCommands(client bot.Client) {
	bot_2.InitialDynamicImport()
	commands.RegisterCommands(client, "", "")
}
