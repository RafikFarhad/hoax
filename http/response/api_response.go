package response

import "github.com/RafikFarhad/hoax/database/model"

func MakeLoginResponse(token string, expiry int64) map[string]interface{} {
	return map[string]interface{}{
		"token":  token,
		"expiry": expiry,
	}
}

func MakeUserResponse(user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"username": user.Username,
		"userInfo": map[string]interface{}{
			"name":   user.UserInfo.Name,
			"avatar": user.UserInfo.Avatar,
		},
	}
}
