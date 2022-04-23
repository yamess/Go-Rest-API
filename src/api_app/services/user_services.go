package services

import (
	"encoding/json"
	"github.com/api_app/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var users []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := append(
		users,
		models.User{
			ID:        uuid.New(),
			FirstName: "Kary",
			LastName:  "Yamess",
			Email:     "kary@yamess.com",
			Password:  "stRoNgpasSworD123",
			IsActive:  true,
			CreatedAt: time.Now(),
		},
		models.User{
			ID:        uuid.New(),
			FirstName: "Bro",
			LastName:  "Brober",
			Email:     "bro@brober.ca",
			Password:  "faKepasswOrd123",
			IsActive:  false,
			CreatedAt: time.Now(),
		})

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		panic("Enable to handle request")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Set the ID and creation date and append to the db
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	users = append(users, user)

	log.Println(user.Email)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	json.NewEncoder(w).Encode(id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User updated")
}
