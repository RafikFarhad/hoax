package model

import (
	"encoding/json"
)

type UserInfo struct {
	*Model
	Name        string `gorm:"not null;size:50" json:"name"`
	Avatar      string `gorm:"size:100" json:"avatar"`
	Institution string `gorm:"size:100" json:"institution"`
	Country     string `gorm:"size:2" json:"country"`
	UserId      uint   `json:"-"`
	// belongsTo
	User *User `json:"-"`
}

func (u *UserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"name":        u.Name,
		"avatar":      u.Avatar,
		"institution": u.Institution,
		"country":     u.Country,
	})
}
