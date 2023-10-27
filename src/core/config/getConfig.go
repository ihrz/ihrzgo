package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Config est la structure de configuration
type Config struct {
	Bot  BotConfig
	API  Api
	Core Core
}

// BotConfig contient la configuration du bot
type BotConfig struct {
	DiscordToken string `toml:"discord_token"`
	ClientID     string `toml:"client_id"`
}

// Api contient la configuration de l'API
type Api struct {
	UseHttps bool `toml:"use_https"`
	HostPort int  `toml:"host_port"`
}

type Core struct {
	Debug               bool   `toml:"debug_mode"`
	BlacklistPictureURL string `toml:"blacklist_picture_url"`
	GuildLogsChannelID  string `toml:"guild_logs_channel_id"`
}

// LoadConfig charge la configuration depuis un fichier TOML
func LoadConfig() Config {
	tomlFilePath := "src/files/config.toml"

	file, err := os.Open(tomlFilePath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier TOML:", err)
		panic(err)
	}
	defer file.Close()
	var config Config

	// Chargez les données depuis le fichier TOML en décodant dans la structure
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		fmt.Println("Erreur lors du chargement du fichier TOML:", err)
		panic(err)
	}

	return config
}

// Configuration publique que vous pouvez utiliser ailleurs
var Main = LoadConfig()
