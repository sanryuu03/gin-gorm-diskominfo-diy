package apilogin

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `validate:"required,min=1" json:"password"`
}
