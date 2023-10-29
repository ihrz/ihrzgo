package db

type SnipeChannelStruc struct {
	SnipeUserInfoTag string
	SnipeUserInfoPp  string
	Snipe            string
	SnipeTimestamp   string
}

func GetLastMessageDeleteHere(guildID *string, channelID *string) SnipeChannelStruc {
	data := SnipeChannelStruc{
		SnipeUserInfoTag: "example",
		SnipeUserInfoPp:  "123",
		Snipe:            "test",
		SnipeTimestamp:   "timestamp",
	}
	return data
}
