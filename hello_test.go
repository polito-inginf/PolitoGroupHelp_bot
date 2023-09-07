package main

import (
	"testing"
)

func TestToTest(t *testing.T) {
	if ToTest(1, 1) != 2 {
		t.Error("Wrong sum!")
	}
}
