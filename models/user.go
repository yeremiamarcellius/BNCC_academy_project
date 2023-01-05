package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique;not null;" json:"username"`
	Password string   `gorm:"not null" json:"password"`
	Email    string   `gorm:"unique;not null" json:"email"`
	Memories []Memory `gorm:"many2many:user_memories;" json:"memories"`
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func UpdateUser(u *User) *User {
	db.Save(&u)
	return u
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	db := db.Where("email=?", email).Find(&user)
	return &user, db.Error
}

func GetUserById(id int64) (*User, error) {
	var user User
	db := db.Where("ID=?", id).Find(&user)
	return &user, db.Error
}

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetAllMemoriesByUser(u *User) []Memory {
	var memories []Memory
	db.Model(&u).Related(&memories, "Memories").Preload("Tags").Preload("User")
	db.Preload("Tags").Find(&memories)
	return memories
}
