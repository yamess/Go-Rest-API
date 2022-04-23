package main

import (
	"fmt"
	"github.com/api_app/constants"
	"github.com/api_app/myserver"
)

func main() {
	fmt.Println("Server starting...")
	fmt.Println("Server started!")
	fmt.Printf("Server running at %s\n", constants.Host)
	myserver.Server()
}
