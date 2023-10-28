package main

import (
	"main/src/Commands/bot"
	"main/src/Commands/fun"
	"main/src/core"
	"main/src/core/db"
)

func main() {
	/*
		Init The Database
	*/
	db.InitDatabase()

	/*
		Fetch All Commands
	*/
	bot.InitialDynamicImport()
	fun.InitialDynamicImport()

	/*
		After all, start the bot
	*/
	core.BotConnect()
}
