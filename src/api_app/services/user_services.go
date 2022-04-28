package services

import (
	"encoding/json"
	"errors"
	"fmt"
	middleware "github.com/api_app/middlewares"
	"github.com/api_app/models"
	"github.com/api_app/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	user := r.Context().Value(middleware.KeyUser{}).(models.User)
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error while hashing the password")
		panic(err.Error())
	}
	result := user.CreateRecord()
	if result.Error != nil {
		log.Print(result.Error)
		http.Error(w, "Enable to insert record", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	// Get db object
	users.ReadRecords()

	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		panic("Enable to handle request")
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	vars := mux.Vars(r)
	id := vars["id"]

	result := user.ReadRecord("id", id)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("Enable to get data from database")
		panic(result.Error)
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		json.NewEncoder(w).Encode("No record found")
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	vars := mux.Vars(r)
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Enable to decode the request body")
		log.Println(err.Error())
		panic(err.Error())
	}

	result := user.UpdateRecord(id)
	if result.Error != nil {
		log.Println("Enable to update data in the database")
		panic(result.Error)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Enable to encode response data")
		panic(err.Error())
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	vars := mux.Vars(r)
	id := vars["id"]

	result := user.DeleteRecord(id)

	if result.Error != nil {
		log.Printf("Enable to delete user with id %s \n", id)
		panic(result.Error)

	} else if result.RowsAffected == 0 {
		response := fmt.Sprintf("No user with id %s exist", id)
		json.NewEncoder(w).Encode(response)
	} else {
		response := fmt.Sprintf("User with id %s deleted", id)
		json.NewEncoder(w).Encode(response)
	}

}
