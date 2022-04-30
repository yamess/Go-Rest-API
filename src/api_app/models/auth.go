package models

import (
	"encoding/json"
	"fmt"
	"github.com/api_app/database"
	"github.com/dgrijalva/jwt-go"
	"io"
	"os"
	"time"
)

type SignIn struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	Token     string `json:"token" validate:"required"`
	TokenType string `json:"token_type" validate:"required"`
}

var EXPIRES = os.Getenv("ACCESS_TOKEN_EXPIRES_MINUTES")
var SecretKEY = os.Getenv("SECRET_KEY")
var JwtKey = []byte(SecretKEY)

func (s *SignIn) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader).Decode(s)
	return e
}

func (s *SignIn) GetUser() (User, error) {
	var user User
	database.PG.Conn.Where("email = ?", s.Email).
		First(&user)
	return user, nil
}

func GenerateJWT(user *User) (string, error) {
	duration, _ := time.ParseDuration(EXPIRES)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["user_id"] = user.ID
	claims["user_email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * duration).Unix()

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		fmt.Printf("Cannot genrate token")
	}
	return tokenString, nil
}
