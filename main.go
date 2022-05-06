package main

import (
	"fmt"
	"net/http"

	"github.com/tomaskul/go-getting-started/controllers"
)

func main() {
	port := 3000
	retries := 3
	_, err := startWebServer(port, retries)
	fmt.Println(port, err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting web server...")
	fmt.Println("   Startup retries: ", numberOfRetries)

	controllers.RegisterControllers()
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	fmt.Println("Web server started. Listening on port: ", port)
	return port, nil
}
