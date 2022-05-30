package models

import (
	"testing"
)

func Test_GetUsers_UninitializedShouldBeNil(t *testing.T) {
	t.Parallel()

	sut := NewUserPersistence()

	if sut.GetUsers() != nil {
		t.Error("Found users when none should be present")
	}
}

func Test_AddUser_SuppliedIdShouldError(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()

	// Act.
	_, err := sut.AddUser(User{7, "FirstName", "Lastname"})

	// Assert.
	if err == nil {
		t.Error("User added when it shouldn't have")
	}
}

func Test_AddUser_ValidInputShouldSucceed(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()

	// Act
	u, err := sut.AddUser(User{0, "Joe", "Bloggs"})

	// Assert.
	if u.ID == 0 && err != nil && len(sut.GetUsers()) != 1 {
		t.Error("User ID not set and/or error returned unexpectedly")
	}
}

func Benchmark_AddUser(b *testing.B) {
	sut := NewUserPersistence()
	for i := 0; i < b.N; i++ {
		sut.AddUser(User{0, "Joe", "Smith"})
	}
}

func Test_GetByUserId_UnknownIdShouldBeNil(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	var inputUsers = [...]*User{{0, "Joe", "Bloggs"},
		{0, "Tim", "Bloggs"},
		{0, "Alice", "Bloggs"}}

	for _, u := range inputUsers {
		sut.AddUser(*u)
	}

	// Act.
	u, err := sut.GetUserByID(7)

	// Assert.
	if (u != User{} || err == nil) {
		t.Error("Returned a user for unknown ID")
	}
}

func Test_GetUserById_ValidIdShouldGetUser(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	var inputUsers = [...]*User{{0, "Joe", "Bloggs"},
		{0, "Tim", "Bloggs"},
		{0, "Alice", "Bloggs"}}

	for _, iu := range inputUsers {
		sut.AddUser(*iu)
	}

	targetId := 2

	// Act.
	u, err := sut.GetUserByID(targetId)

	// Assert.
	if u.ID != targetId || err != nil || u.FirstName != "Tim" {
		t.Errorf("Returned unknown result, u.FirstName: %s", u.FirstName)
	}
}

func Test_UpdateUser_IdTooLowShouldError(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	sut.AddUser(User{0, "Joe", "Bloggs"})

	// Act.
	u, err := sut.UpdateUser(User{0, "Simon", "Bloggs"})

	// Assert.
	if (u != User{} || err == nil) {
		t.Error("Failed to validate user IDs below 1")
	}
}

func Test_UpdateUser_UserFoundShouldUpdateValues(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	u, _ := sut.AddUser(User{0, "Joe", "Bloggs"})
	newLastName := "Jamerson"

	// Act.
	u.LastName = newLastName
	u2, err := sut.UpdateUser(u)

	// Assert.
	if (u2 == User{} || err != nil) {
		t.Error("Failed to update user")
	}
	if u.FirstName != u2.FirstName || u2.LastName != newLastName ||
		u.ID != u2.ID {
		t.Error("Update failed")
	}
}

func Test_UpdateUser_IdNotInRangeShouldError(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	u1, _ := sut.AddUser(User{0, "Alice", "Bloggs"})
	sut.AddUser(User{0, "Bob", "Bloggs"})
	u3, _ := sut.AddUser(User{0, "Cindy", "Bloggs"})

	// Act.
	u1.ID = u3.ID + 3
	u1.FirstName = "Barbara"
	u, err := sut.UpdateUser(u1)

	// Assert.
	if (u != User{} || err == nil) {
		t.Error("Failed to validate user IDs below 1")
	}
}

func Test_RemoveUserById_IdNotFoundShouldError(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	inputUser, _ := sut.AddUser(User{0, "Alice", "Bloggs"})

	// Act.
	err := sut.RemoveUserById(inputUser.ID + 78)

	// Assert.
	if err == nil {
		t.Error("Removed user with non-existent ID.")
	}
}

func Test_RemoveUserById_RemovedUserShouldReturnNil(t *testing.T) {
	t.Parallel()

	// Arrange.
	sut := NewUserPersistence()
	user, _ := sut.AddUser(User{0, "Alice", "Bloggs"})

	// Act.
	err := sut.RemoveUserById(user.ID)

	// Assert.
	u, err2 := sut.GetUserByID(user.ID)

	if (err != nil || u != User{} || err2 == nil) {
		t.Error("Failed to remove user by ID.")
	}
}
