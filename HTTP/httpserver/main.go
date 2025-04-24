package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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

var response = map[string]any{}
var usersDb = []User{
	{
		Id:      "1",
		Name:    "Tuhtaboy",
		Age:     30,
		Address: "123 Main St"},
	{
		Id:      "2",
		Name:    "Ali",
		Age:     25,
		Address: "456 Elm St"},
	{
		Id:      "3",
		Name:    "Bob",
		Age:     35,
		Address: "789 Oak St"},
	{
		Id:      "4",
		Name:    "Alice",
		Age:     28,
		Address: "101 Pine St"},
}

func users(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		createUser(w, r)
		return
	} else if method == "PUT" {
		updateUser(w, r)
		return
	}

	response["status"] = "success"
	response["statusCode"] = http.StatusOK

	//usersjson, _ := json.Marshal(usersDb)
	response["data"] = usersDb
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func singleUsert(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	for _, u := range usersDb {
		if u.Id == id {

			response["data"] = u
			response["status"] = "success"
			response["statusCode"] = http.StatusOK

			json.NewEncoder(w).Encode(response)

			return
		}
	}
	response["status"] = "Not Found"
	response["statusCode"] = http.StatusNotFound
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(response)

}
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser CreateUserDTO
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.Id = string(time.Now().Unix())
	usersDb = append(usersDb, User(newUser))
	response["status"] = "Created"
	response["statusCode"] = http.StatusCreated
	json.NewEncoder(w).Encode(response)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}
func updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	id := mux.Vars(r)["id"]

	for i, u := range usersDb {
		if u.Id == id {
			usersDb[i] = updatedUser
			response["status"] = "Updated"
			response["statusCode"] = http.StatusOK
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response["status"] = "Not Found"
	response["statusCode"] = http.StatusNotFound
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(response)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, u := range usersDb {
		if u.Id == id {
			usersDb = append(usersDb[:i], usersDb[i+1:]...)
			response["status"] = "Deleted"
			response["statusCode"] = http.StatusOK
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	response["status"] = "Not Found"
	response["statusCode"] = http.StatusNotFound
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Starting HTTP server on port 8080...")

	r := mux.NewRouter()
	r.HandleFunc("/users", users)
	r.HandleFunc("/users/{id}", singleUsert)
	r.HandleFunc("/users", createUser)
	r.HandleFunc("/users/{id}", updateUser)
	r.HandleFunc("/users/{id}", deleteUser)

	{
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal(err)
		}

	}
}
