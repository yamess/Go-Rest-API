package main

import (
	"fmt"
	"github.com/api_app/myserver"
)

func main() {
	fmt.Println("Server starting...")
	fmt.Println("Server started at 0.0.0.0:8081")
	myserver.Server()
}
