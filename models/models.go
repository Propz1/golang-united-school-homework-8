package models

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Age   int `json:"age"`
}