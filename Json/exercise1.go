package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {

	user := User{
		Name:  "Sardor",
		Age:   35,
		Email: "anjonlik@yahoo.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
