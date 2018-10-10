package main

import (
	"testing"
)

func TestDiamond(t *testing.T) {

	result := Diamond('A')

	expected := "A"

	if result != expected {
		t.Errorf("Result %s != expected %s", result, expected)
	}

}
