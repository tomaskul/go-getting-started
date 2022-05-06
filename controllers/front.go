package controllers

import "net/http"

// Registers all controllers supported by this web service.
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
