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

func TestReverseOdd(t *testing.T) {

	result := reverse("words")

	expected := "sdrow"

	if result != expected {
		t.Errorf("'%s' != '%s'", result, expected)
	}
}

func TestReverseAcento(t *testing.T) {

	result := reverse("até")

	expected := "éta"

	if result != expected {
		t.Errorf("'%s' != '%s'", result, expected)
	}
}

func TestReverseCharacter(t *testing.T) {

	result := reverse("até+")

	expected := "+éta"

	if result != expected {
		t.Errorf("'%s' != '%s'", result, expected)
	}
}
