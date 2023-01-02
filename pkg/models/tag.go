package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	// ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (t *Tag) CreateTag() *Tag {
	db.NewRecord(t)
	db.Create(&t)
	return t
}

func GetAllTags() []Tag {
	var Tags []Tag
	db.Find(&Tags)
	return Tags
}

func GetTagById(id uint) (*Tag, *gorm.DB) {
	var tag Tag
	db := db.Where("ID=?", id).Find(&tag)
	return &tag, db
}

func DeleteTag(id uint) Tag {
	var tag Tag
	db.Where("ID=?", id).Delete(&tag)
	return tag
}
