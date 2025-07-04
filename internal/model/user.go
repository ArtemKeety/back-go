package model

type UserData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRequest struct {
	UserData
	Email string `json:"email"`
}

type User struct {
	Guid string `json:"guid"`
	UserRequest
}
