package model

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
	UserName  string `json:"username"`
	Location  string `json:"location"`
	FreeSpace string `json:"free_space"`
}
