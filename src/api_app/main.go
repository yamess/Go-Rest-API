package main

import (
	"fmt"
	"github.com/api_app/constants"
	db "github.com/api_app/database"
	"github.com/api_app/models"
	"github.com/api_app/myserver"
)

func init() {

}

func main() {
	fmt.Println("Server starting...")
	fmt.Println("Server started!")
	fmt.Printf("Server running at %s\n", constants.Host)

	db.AutoMigrate(&models.User{})
	//db.AutoMigrate()

	myserver.Server()
}
