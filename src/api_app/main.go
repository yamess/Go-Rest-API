package main

import (
	"fmt"
	"github.com/api_app/constants"
	db "github.com/api_app/database"
	"github.com/api_app/models"
	"github.com/api_app/myserver"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading the environment file")
	}
	fmt.Println("Server starting...")
	fmt.Println("Server started!")
	fmt.Printf("Server running at %s\n", constants.Host)

	db.AutoMigrate(&models.User{})

	myserver.Server()
}
