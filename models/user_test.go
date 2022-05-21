package models

import (
	"testing"
)

func Test_GetUsers_UninitializedShouldBeNil(t *testing.T) {
	if GetUsers() != nil {
		t.Error("Found users when none should be present")
	}
}

func Test_AddUser_SuppliedIdShouldError(t *testing.T) {
	t.Parallel()
	_, err := AddUser(User{7, "FirstName", "Lastname"})
	if err == nil {
		t.Error("User added when it shouldn't have")
	}
}

func Test_AddUser_ValidInputShouldSucceed(t *testing.T) {
	t.Parallel()
	u, err := AddUser(User{0, "Joe", "Bloggs"})

	if u.ID == 0 && err != nil && len(GetUsers()) != 1 {
		t.Error("User ID not set and/or error returned unexpectedly")
	}
}

func Test_GetByUserId_UnknownIdShouldBeNil(t *testing.T) {
	t.Parallel()
	// Arrange.
	var inputUsers = [...]*User{{0, "Joe", "Bloggs"},
		{0, "Tim", "Bloggs"},
		{0, "Alice", "Bloggs"}}

	for _, u := range inputUsers {
		AddUser(*u)
	}

	// Act.
	u, err := GetUserByID(7)

	// Assert.
	if (u != User{} || err == nil) {
		t.Error("Returned a user for unknown ID")
	}
}

func Test_GetUserById_ValidIdShouldGetUser(t *testing.T) {
	t.Parallel()
	// Arrange.
	var inputUsers = [...]*User{{0, "Joe", "Bloggs"},
		{0, "Tim", "Bloggs"},
		{0, "Alice", "Bloggs"}}

	for _, iu := range inputUsers {
		AddUser(*iu)
	}

	targetId := 2

	// Act.
	u, err := GetUserByID(targetId)

	// Assert.
	if u.ID != targetId || err != nil || u.FirstName != "Tim" {
		t.Errorf("Returned unknown result, u.FirstName: %s", u.FirstName)
	}
}
