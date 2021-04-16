package model

import (
	"encoding/json"
	"github.com/RafikFarhad/hoax/database"
)

type User struct {
	*Model
	Username string `gorm:"uniqueIndex;not null;size:30" json:"username"`
	Email    string `gorm:"uniqueIndex;size:50" json:"email"`
	Password string `gorm:"not null;size:50" json:"-"`
	// hasOne
	UserInfo *UserInfo `gorm:"foreignKey:UserId"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":       u.Id,
		"username": u.Username,
	})
}

func GetUserById(user *User, id uint, preloads ...string) error {
	tx := database.AppDb
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	result := tx.First(user, id)
	return result.Error
}

func GetUserByUsername(user *User, username string) error {
	return getUserByKeyAndValue(user, "username", username)
}

func getUserByKeyAndValue(user *User, key string, value string) error {
	result := database.AppDb.Where(key+" = ?", value).First(user)
	return result.Error
}
