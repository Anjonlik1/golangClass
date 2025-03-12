package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	jsonData := `{ "first_name": "Alice", "last_name": "Smith" }`

	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Printf("user: %+v\n", user)

}
