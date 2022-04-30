package services

import (
	"encoding/json"
	"github.com/api_app/models"
	"github.com/api_app/utils"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var signIn models.SignIn
	//var user models.Users

	signIn.FromJSON(r.Body)
	user, _ := signIn.GetUser()

	isPwdValid := utils.CheckPasswordHash(signIn.Password, user.Password)
	if !isPwdValid {
		http.Error(w, "Cannot sign in", http.StatusBadRequest)
		return
	}
	token, _ := models.GenerateJWT(&user)
	auth := models.Auth{Token: token, TokenType: "bearer"}

	json.NewEncoder(w).Encode(auth)

}
