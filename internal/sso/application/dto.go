package application

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type LoginResponse struct {
	AccessToken  Token `json:"token"`
	RefreshToken Token `json:"refresh_token"`
}
