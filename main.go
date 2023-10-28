package main

import (
	commands "main/src/Commands"
	"main/src/core"
	"main/src/core/db"
)

func main() {
	db.InitDatabase()
	commands.InitialDynamicImport()
	core.BotConnect()
}
