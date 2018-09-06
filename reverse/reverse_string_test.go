package main

import (
	"testing"
)

func TestReverse(t *testing.T) {

	result := reverse("word")

	expected := "drow"

	if result != expected {
		t.Errorf("'%s' != '%s'", result, expected)
	}
}
