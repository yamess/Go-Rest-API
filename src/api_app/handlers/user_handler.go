package handlers

import (
	"encoding/json"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("All users in db")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("New user created")
}

func SingleUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("A single user data")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User updated")
}
