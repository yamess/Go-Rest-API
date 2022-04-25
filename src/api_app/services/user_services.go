package services

import (
	"encoding/json"
	"fmt"
	"github.com/api_app/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	result := user.CreateRecord()

	if result.RowsAffected > 0 {
		log.Println("New User created")

		err = json.NewEncoder(w).Encode(&user)
		if err != nil {
			panic(err)
		}
	} else {
		err = json.NewEncoder(w).Encode("Enable to insert record")
		if err != nil {
			panic(err)
		}
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

	if result.Error != nil {
		log.Println("Enable to get data from database")
		panic(result.Error)
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Enable to encode response")
		panic(err.Error())
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
