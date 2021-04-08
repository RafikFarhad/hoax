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

func (u *User) WithUserInfo() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.Id,
		"userInfo": u.UserInfo,
	}
}

func GetUserById(id uint, preloads ...string) (*User, error) {
	user := &User{}
	tx := database.AppDb
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}
	result := tx.First(user, id)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, result.Error
}

func getByKeyAndValue(key string, value string) (*User, error) {
	user := &User{}
	result := database.AppDb.Where(key+" = ?", value).First(user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, result.Error
}

func GetUserByUsername(username string) (*User, error) {
	return getByKeyAndValue("username", username)
}

func GetUserByEmail(email string) (*User, error) {
	return getByKeyAndValue("email", email)
}
