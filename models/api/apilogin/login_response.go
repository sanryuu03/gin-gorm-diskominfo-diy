package apilogin

type LoginResponse struct {
	Email           string `json:"email"`
	AccessTokenKey  string `json:"access_token_key"`
	RefreshTokenKey string `json:"refresh_token_key"`
}
