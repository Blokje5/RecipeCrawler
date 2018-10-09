package entity

import "testing"

func TrivialTest(t *testing.T) {
	total := 1+1
	if total != 2 {
		t.Errorf("failed")
	}
}