package auth

type RegisterModel struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
