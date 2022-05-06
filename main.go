package main

import (
	"fmt"

	"github.com/tomaskul/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)
}
