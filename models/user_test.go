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
