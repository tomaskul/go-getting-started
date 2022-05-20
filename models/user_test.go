package models

import (
	"testing"
)

func Test_GetUsers_UninitializedShouldBeNil(t *testing.T) {
	if GetUsers() != nil {
		t.Error("Found users when none should be present")
	}
}

func Test_AddUser_ErrorOnSuppliedId(t *testing.T) {
	var u = User{7, "FirstName", "Lastname"}
	_, err := AddUser(u)
	if err == nil {
		t.Error("User added when it shouldn't have")
	}
}

func Test_AddUser_Success(t *testing.T) {
	var u = User{0, "Joe", "Bloggs"}
	u2, err := AddUser(u)

	if u2.ID == 0 && err != nil && len(GetUsers()) != 1 {
		t.Error("User ID not set and/or error returned unexpectedly")
	}
}
