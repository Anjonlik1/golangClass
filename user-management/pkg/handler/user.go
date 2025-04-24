package handler

import (
	"encoding/json"
	"net/http"
	"user-management/model"

	"github.com/gorilla/mux"
)

func (h *Handler) users(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Set("Content-type", "application/json")

	if method == "POST" {
		h.createUser(w, r)
		return
	} else if method == "PUT" {
		// updateUser(w, r)
		return
	}

	response["status"] = "Success"
	response["statusCode"] = http.StatusOK

	response["data"] = h.services.GetUsers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) singleUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	w.Header().Set("Content-type", "application/json")
	if method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := mux.Vars(r)["id"]
	user, ok := h.services.GetSingleUser(id)

	if ok {
		response["data"] = user
		response["status"] = "Success"
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
	method := r.Method
	if method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-type", "application/json")
	var newUser model.CreateUserDTO
	json.NewDecoder(r.Body).Decode(&newUser)
	// Service for creation of new user
	h.services.User.CreateUser(newUser)

	response["status"] = "Created"
	response["statusCode"] = http.StatusCreated
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
