package services

import (
	"bytes"
	"encoding/json"
	"github.com/api_app/models"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var user = models.User{
	FirstName: "Kary",
	LastName:  "Yamess",
	Email:     "kary@yamess.com",
	Password:  "testerPwd123",
	IsActive:  true,
}

func TestCreateUser(t *testing.T) {

	requestBody, err := json.Marshal(map[string]any{
		"first_name": "Abou",
		"last_name":  "yamess",
	})

	if err != nil {
		t.Errorf("Error serializing request body")
	}
	response, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error while reading the response body")
	}
	log.Println(string(body))
}

func TestUserValidation(t *testing.T) {
	user := &models.User{
		FirstName: "Kary",
		LastName:  "Yamess",
		Email:     "yameogo.boukari@gmail.com",
		Password:  "stRoNgsSworD123",
		IsActive:  true,
	}

	err := user.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUsers(t *testing.T) {

}

func TestGetUserByID(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}

//func TestCreateProduct(t *testing.T) {
//
//	clearTable()
//
//	var jsonStr = []byte(`{"name":"test product", "price": 11.22}`)
//	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))
//	req.Header.Set("Content-Type", "application/json")
//
//	response := executeRequest(req)
//	checkResponseCode(t, http.StatusCreated, response.Code)
//
//	var m map[string]interface{}
//	json.Unmarshal(response.Body.Bytes(), &m)
//
//	if m["name"] != "test product" {
//		t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
//	}
//
//	if m["price"] != 11.22 {
//		t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
//	}
//
//	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
//	// floats, when the target is a map[string]interface{}
//	if m["id"] != 1.0 {
//		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
//	}
//}
