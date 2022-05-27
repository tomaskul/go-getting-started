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

func NotTest_UpdateUser_UserFoundShouldUpdateValues(t *testing.T) {

}

func NotTest_UpdateUser_IdNotInRangeShouldError(t *testing.T) {

}

func NotTest_RemoveUserById_IdNotFoundShouldError(t *testing.T) {

}

func NotTest_RemoveUserById_RemovedUserShouldReturnNil(t *testing.T) {

}
