package model

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
