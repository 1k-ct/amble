package model

import "time"

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
	UserName  string `json:"username"`
	Location  string `json:"location"`
	FreeSpace string `json:"free_space"`
}
type Flowers struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name"`
	ToFlowID  string    `json:"to_flow_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
