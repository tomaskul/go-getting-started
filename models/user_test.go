package models

import (
	"testing"
)

func Test_BasicSubtraction(t *testing.T) {
	if 10-5 != 5 {
		t.Error("Failed to subtract correctly")
	}
}
