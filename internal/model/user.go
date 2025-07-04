package model

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Guid string `json:"guid"`
	UserRequest
}
