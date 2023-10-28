package events

import (
	commands "main/src"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	commands.RegisterCommands(s, s.State.User.ID, "")
	changeStatus(s)
}

func changeStatus(s *discordgo.Session) {
	quotes := []string{
		"discord.gg/ihorizon",
		"https://ihorizon.me",
		"iHorizon x ElektraBots <3",
		"Did you know you can have your own iHorizon? For really cheap??",
		"Our goal is to make the internet simpler!",
		"My goal is to make internet so simple that my own mother can use it!",
		"280K USERS !? 🥳🥳🥳",
		"It's not 250k anymore it's 280k 😎😎😎😎",
		"trusted by big servers 😎",
		"Nah men I'm not getting paid enough to manage 280K users...",
		"Never gonna give you up...BRO YOU'VE BEEN RICK ROLLED BY A BOT",
		"I have a youtube channel!",
		"Youtube, X (twitter), only****, what's next?",
		"Github is basically onlyfan for code, so I have an onlyfan 😎",
		"My owner doesn't use tiktok...I INSTALLED IT BEHIND HER BACK",
		"I removed my own database (going insane) 😎😎😎",
		"I can code myself (Not a joke)",
		"I BROKED MY CODE HELP ME",
		"What is a database? Do I really need one?",
		"20 bucks for my token",
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]

	s.UpdateGameStatus(0, randomQuote)
}
