package model

import (
	"encoding/json"
	"github.com/RafikFarhad/hoax/app"
)

type User struct {
	*Model
	Username string `gorm:"uniqueIndex;not null;size:30" json:"username"`
	Email    string `gorm:"uniqueIndex;size:50" json:"email"`
	Password string `gorm:"not null;size:50" json:"-"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"user":  u.Username,
		"email": u.Email,
	})
}

func GetUserById(id int) (*User, error) {
	user := &User{}
	result := app.App.Db.First(user, id)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, result.Error
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	result := app.App.Db.Where("username = ?", username).First(user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, result.Error
}
