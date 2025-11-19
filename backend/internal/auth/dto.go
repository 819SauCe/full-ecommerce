package auth

import "errors"

var (
	ErrInvalidFirstName   = errors.New("invalid_first_name")
	ErrInvalidLastName    = errors.New("invalid_last_name")
	ErrInvalidEmail       = errors.New("invalid_email")
	ErrEmailAlreadyUsed   = errors.New("email_already_used")
	ErrInvalidPassword    = errors.New("invalid_password")
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type MeResponse struct {
	Status     string `json:"status"`
	ID         any    `json:"id"`
	ProfileImg any    `json:"profile_img"`
	FirstName  any    `json:"first_name"`
	LastName   any    `json:"last_name"`
	Email      any    `json:"email"`
	Role       any    `json:"role"`
	Exp        any    `json:"exp"`
	Iat        any    `json:"iat"`
}
