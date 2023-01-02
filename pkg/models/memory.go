package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Memory struct {
	gorm.Model
	// ID    uint      `gorm:"primary_key" json:"id"`
	Title string    `gorm:"unique;not null;" json:"title"`
	Image string    `gorm:"not null" json:"image"`
	Date  time.Time `gorm:"not null" json:"date"`
	Tags  []Tag     `gorm:"many2many:memory_tags;" json:"tags"`
	Desc  string    `gorm:"not null" json:"desc"`
}

func (m *Memory) CreateMemory() *Memory {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMemories() []Memory {
	var Memories []Memory
	db.Find(&Memories)
	return Memories
}

func GetMemoryById(id int64) (*Memory, *gorm.DB) {
	var memory Memory
	db := db.Where("ID=?", id).Find(&memory)
	return &memory, db
}

func UpdateMemory(m *Memory) *Memory {
	db.Save(&m)
	return m
}

func DeleteMemory(id int64) Memory {
	var memory Memory
	db.Where("ID=?", id).Delete(&memory)
	return memory
}
