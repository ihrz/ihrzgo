package main

import (
	commands "main/src/Commands"
	"main/src/core"
)

func main() {
	core.InitDatabase()
	commands.InitialDynamicImport()
	core.BotConnect()
}
