package handler

import (
	"encoding/json"
	"net/http"
	"user-management/model"

	"github.com/gorilla/mux"
)

func (h *Handler) users(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		h.createUser(w, r)
		return
	} else if method == "PUT" {
		h.updateUser(w, r)
		return
	} else if method == "DELETE" {
		h.deleteUser(w, r)
		return
	}
	response["status"] = "success"
	response["statusCode"] = http.StatusOK

	//usersjson, _ := json.Marshal(usersDb)
	response["data"] = h.services.User.GetUser()
	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func (h *Handler) singleUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	user, ok := h.services.User.GetUserById(id)
	if ok {
		response["data"] = user
		response["status"] = "success"
		response["statusCode"] = http.StatusOK

		json.NewEncoder(w).Encode(response)

		return
	}

	response["status"] = "Not Found"
	response["statusCode"] = http.StatusNotFound
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(response)

}
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.CreateUserDTO
	json.NewDecoder(r.Body).Decode(&newUser)

	h.services.User.CreateUser(newUser)

	response["status"] = "Created"
	response["statusCode"] = http.StatusCreated
	json.NewEncoder(w).Encode(response)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

}
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
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
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
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
