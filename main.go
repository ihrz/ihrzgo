package main

import (
	"context"
	"log/slog"
	"main/config"
	"main/core"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/snowflake/v2"
)

var (
	token   = config.Main.Bot.DiscordToken
	guildID = snowflake.GetEnv("disgo_guild_id")

	// commands = []discord.ApplicationCommandCreate{
	// 	discord.SlashCommandCreate{
	// 		Name:        "say",
	// 		Description: "says what you say",
	// 		Options: []discord.ApplicationCommandOption{
	// 			discord.ApplicationCommandOptionString{
	// 				Name:        "message",
	// 				Description: "What to say",
	// 				Required:    true,
	// 			},
	// 			discord.ApplicationCommandOptionBool{
	// 				Name:        "ephemeral",
	// 				Description: "If the response should only be visible to you",
	// 				Required:    true,
	// 			},
	// 		},
	// 	},
	// }
)

func main() {
	slog.Info("starting example...")
	slog.Info("disgo version", slog.String("version", disgo.Version))

	client, err := disgo.New(token,
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuilds, gateway.IntentGuildMessages, gateway.IntentDirectMessages, gateway.IntentMessageContent)),
		bot.WithEventListenerFunc(commandListener),
	)
	core.SyncCommands(client)

	if err != nil {
		slog.Error("error while building bot", slog.Any("err", err))
		return
	}

	defer client.Close(context.TODO())

	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error("error while connecting to gateway", slog.Any("err", err))
		return
	}

	slog.Info("example is now running. Press CTRL-C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s

}

func commandListener(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()
	if data.CommandName() == "say" {
		err := event.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(data.String("message")).
			SetEphemeral(data.Bool("ephemeral")).
			Build(),
		)
		if err != nil {
			slog.Error("error on sending response", slog.Any("err", err))
		}
	}
}
