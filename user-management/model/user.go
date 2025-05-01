package model

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
type CreateUserDTO struct {
	Id      string
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
