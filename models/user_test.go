package models_test

import (
	"testing"

	"github.com/tomaskul/go-getting-started/models"
)

func Test_GetUsers_UninitializedShouldBeNil(t *testing.T) {
	if models.GetUsers() != nil {
		t.Error("Found users when none should be present")
	}
}

func Test_AddUser(t *testing.T) {

	t.Run("SuppliedIdShouldError", func(t *testing.T) {
		// Act.
		_, err := models.AddUser(models.User{7, "FirstName", "Lastname"})

		// Assert.
		if err == nil {
			t.Error("User added when it shouldn't have")
		}
	})

	t.Run("ValidInputShouldSucceed", func(t *testing.T) {
		// Act
		u, err := models.AddUser(models.User{0, "Joe", "Bloggs"})

		// Assert.
		if u.ID == 0 && u.ID != 1 && err != nil && len(models.GetUsers()) != 1 {
			t.Error("User ID not set and/or error returned unexpectedly")
		}
	})

	// Tear-down.
	models.Reset()
}

func Benchmark_AddUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		models.AddUser(models.User{0, "Joe", "Smith"})
	}
}

func Test_GetByUserId_UnknownIdShouldBeNil(t *testing.T) {
	// Arrange.
	models.AddUser(models.User{0, "Jay", "Sloggs"})
	models.AddUser(models.User{0, "Tim", "Ploggs"})
	models.AddUser(models.User{0, "Alice", "Wloggs"})

	t.Run("UnknownIdShouldBeNil", func(t *testing.T) {
		// Act.
		u, err := models.GetUserByID(7)

		// Assert.
		if (u != models.User{} || err == nil) {
			t.Error("Returned a user for unknown ID")
		}
	})

	t.Run("ValidIdShouldGetUser", func(t *testing.T) {
		targetId := 2
		// Act.
		u, err := models.GetUserByID(targetId)

		// Assert.
		if u.ID != targetId || err != nil || u.FirstName != "Tim" {
			t.Errorf("Returned unknown result, u: %v", u)
		}
	})

	// Tear-down
	models.Reset()
}

func Test_UpdateUser(t *testing.T) {
	// Arrange.
	u, _ := models.AddUser(models.User{0, "Joe", "Logs"})
	newLastName := "Jamerson"

	models.AddUser(models.User{0, "Bob", "Bloggs"})
	u3, _ := models.AddUser(models.User{0, "Cindy", "Bloggs"})

	t.Run("IdTooLowShouldError", func(t *testing.T) {
		// Act.
		u, err := models.UpdateUser(models.User{0, "Simon", "Bloggs"})

		// Assert.
		if (u != models.User{} || err == nil) {
			t.Error("Failed to validate user IDs below 1")
		}
	})

	t.Run("UserFoundShouldUpdateValues", func(t *testing.T) {
		// Act.
		u.LastName = newLastName
		u2, err := models.UpdateUser(u)

		// Assert.
		if (u2 == models.User{} || err != nil) {
			t.Error("Failed to update user")
		}
		if u.FirstName != u2.FirstName || u2.LastName != newLastName ||
			u.ID != u2.ID {
			t.Error("Update failed")
		}
	})

	t.Run("IdNotInRangeShouldError", func(t *testing.T) {
		// Act.
		u.ID = u3.ID + 3
		u.FirstName = "Barbara"
		u, err := models.UpdateUser(u)

		// Assert.
		if (u != models.User{} || err == nil) {
			t.Error("Failed to validate user IDs below 1")
		}
	})

	// Tear-down
	models.Reset()
}

func Test_RemoveUserById(t *testing.T) {
	// Arrange.
	inputUser, _ := models.AddUser(models.User{0, "Alice", "Bloggs"})

	t.Run("IdNotFoundShouldError", func(t *testing.T) {
		// Act.
		err := models.RemoveUserById(inputUser.ID + 78)

		// Assert.
		if err == nil {
			t.Error("Removed user with non-existent ID.")
		}
	})

	t.Run("RemovedUserShouldReturnNil", func(t *testing.T) {
		// Act.
		err := models.RemoveUserById(inputUser.ID)

		// Assert.
		u, err2 := models.GetUserByID(inputUser.ID)

		if (err != nil || u != models.User{} || err2 == nil) {
			t.Error("Failed to remove user by ID.")
		}
	})

	// Tear-down
	models.Reset()
}
