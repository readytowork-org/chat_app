package dtos

// Request body data to authenticate user with jwt-auth
type JWTLoginRequestData struct {
	Phone    string `json:"Phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
