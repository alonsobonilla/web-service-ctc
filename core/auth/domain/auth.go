package authdomain

type Auth struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	TipoLogin string `json:"tipo_login"`
}

type AuthResponse struct {
	Jwt string `json:"jwt"`
}
