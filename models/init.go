package models

import (
	"bncc/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.DropTableIfExists(&Memory{}, &Tag{}, &User{})
	db.AutoMigrate(&Memory{}, &Tag{}, &User{})
}
