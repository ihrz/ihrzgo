package db

import (
	"main/src/core"

	"gorm.io/gorm"
)

var db = core.DataBase()

func GetAllowList(str *string) {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}
	var product Product

	db.First(&product, "code = ?", "D42")
}

// 	baseData := getAllowList(message.GuildID)

// 	if baseData == nil {
// 		setAllowList(message.GuildID, &AllowListConfig{
// 			Enable: false,
// 			List: map[string]AllowListEntry{
// 				message.Guild.OwnerID: {Allowed: true},
// 			},
// 		})
// 	}
