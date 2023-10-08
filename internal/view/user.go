package view

type UserCreateView struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginView struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserView struct {
}

type LoginResponseView struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
