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
	} else if method == "GET" {
		response["status"] = "success"
		response["statusCode"] = http.StatusOK

		response["data"] = h.services.User.GetUsers()
		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}

}

func (h *Handler) singleUser(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	user, err := h.services.User.GetSingleUser(id)
	if err != nil {
		response["status"] = "Not Found"
		response["statusCode"] = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
	} else {
		response["data"] = user
		response["status"] = "success"
		response["statusCode"] = http.StatusOK
	}
	json.NewEncoder(w).Encode(response)
}
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.CreateUserDTO
	json.NewDecoder(r.Body).Decode(&newUser)

	h.services.User.CreateUser(newUser)

	response["status"] = "Created"
	response["statusCode"] = http.StatusCreated
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var updatedUser model.User
	json.NewDecoder(r.Body).Decode(&updatedUser)
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	founded := h.services.UpdateUser(id, updatedUser)
	if !founded {
		response["status"] = "Not Found"
		response["statusCode"] = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
	} else {
		response["status"] = "Ok"
		response["statusCode"] = http.StatusOK
	}

	json.NewEncoder(w).Encode(response)
}
func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	err := h.services.DeleteUser(id)
	if err != nil {
		response["status"] = "Not Found"
		response["statusCode"] = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
	} else {
		response["status"] = "Ok"
		response["statusCode"] = http.StatusOK
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}
