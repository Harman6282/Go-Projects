package models

type User struct {
	UserId   string `json:"user_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int32  `json:"age,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
