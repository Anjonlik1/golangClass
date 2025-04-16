package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
	{
		data: [],
		status: "Success" | "Error"
		statusCode: http.StatusCode
	}
*/
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
	{"1", "John", 23, "New-York"},
	{"2", "Sarah", 22, "Washington"},
	{"3", "Emmy", 20, "Warsaw"},
	{"4", "Sardor", 25, "Tashkent"},
}

func users(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Set("Content-type", "application/json")

	if method == "POST" {
		createUser(w, r)
		return
	} else if method == "PUT" {
		// updateUser(w, r)
		return
	}

	response["status"] = "Success"
	response["statusCode"] = http.StatusOK

	response["data"] = usersDb

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func singleUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Set("Content-type", "application/json")
	if method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := mux.Vars(r)["id"]

	for _, u := range usersDb {
		if u.Id == id {
			response["data"] = u
			response["status"] = "Success"
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
	method := r.Method
	if method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	var newUser CreateUserDTO
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.Id = string(time.Now().Unix())
	usersDb = append(usersDb, User(newUser))
	response["status"] = "Created"
	response["statusCode"] = http.StatusCreated
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Server is running on http://localhost:8090")

	r := mux.NewRouter()

	r.HandleFunc("/users", users)
	r.HandleFunc("/users/{id}", singleUser)
	r.HandleFunc("/users", createUser)

	if err := http.ListenAndServe(":8090", r); err != nil {
		log.Fatal(err)
	}
}
