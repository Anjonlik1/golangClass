package main

import "fmt"

type User struct {
	Name string
}

func (u User) ChangeName(newName string) {
	u.Name = newName
}
func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	u := User{Name: "Alice"}

	u.ChangeName("Bob")
	fmt.Println("value change:", u.Name) // Won't modify original

	u.UpdateName("Charlie") // Will modify original
	fmt.Println("pointer change:", u.Name)
}
