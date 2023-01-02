package models

import (
	"github.com/jinzhu/gorm"
	"github.com/youhane/bncc_academy_final/pkg/config"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.DropTableIfExists(&Memory{}, &Tag{}, &User{})
	db.AutoMigrate(&Memory{}, &Tag{}, &User{})
}
