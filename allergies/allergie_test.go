package main

import (
	"strings"
	"testing"
)

func TestAllergie(t *testing.T) {

	result := allergie(34)
	expected := "peanuts chocolate "

	if result != expected {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}

func TestAllergieBig(t *testing.T) {

	result := allergie(257)
	expected := "eggs "

	if result != expected {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}

func TestAllergie0(t *testing.T) {

	result := allergie(0)
	expected := ""

	if strings.Compare(result, expected) != 0 {
		t.Errorf("Result: %s, expected: %s", result, expected)
	}
}
