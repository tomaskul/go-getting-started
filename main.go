package main

import (
	"fmt"

	"github.com/tomaskul/go-getting-started/models"
)

func main() {
	startWebServer()

	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)
}

func startWebServer() {
	fmt.Println("Starting web server...")
	fmt.Println("Web server started.")
}
