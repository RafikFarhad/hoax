package response

type LoginResponse struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}
