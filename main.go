package main

import (
	commands "main/src/Commands"
	"main/src/core"
)

func main() {
	core.DataBase()
	commands.InitialDynamicImport()
	// core.EventManager()
	core.BotConnect()
}
