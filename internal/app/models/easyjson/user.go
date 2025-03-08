package easyjson

//go:generate easyjson -all user.go

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
