package main

import (
	commands "main/src/Commands"
	"main/src/core"
)

func main() {
	commands.InitialDynamicImport()
	core.EventManager()

	core.BotConnect()
}
