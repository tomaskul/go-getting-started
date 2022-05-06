package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	retries := 3
	_, err := startWebServer(port, retries)
	fmt.Println(port, err)

	/*
		u := models.User{
			ID:        2,
			FirstName: "Tricia",
			LastName:  "McMillan",
		}

		fmt.Println(u)*/
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting web server...")
	fmt.Println("   Startup retries: ", numberOfRetries)

	return -1, errors.New("something went wrong")

	//fmt.Println("Web server started. Listening on port: ", port)
	//return port, nil
}
