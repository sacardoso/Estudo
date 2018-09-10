package main

import (
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {

	A := "1, 2, 3"
	B := "1, 2, 3"

	result := Sublist(A, B)

	expected := "A is equal to B"

	if strings.Contains(result, expected) == false {
		t.Errorf("Resultado %s != %s", result, expected)
	}
}

func TestSublist(t *testing.T) {

	A := "1, 2"
	B := "1, 2, 3"

	result := Sublist(A, B)

	expected := "A is a sublist of B"

	if strings.Contains(result, expected) == false {
		t.Errorf("Resultado %s != %s", result, expected)
	}
}

func TestSuperlist(t *testing.T) {

	A := "1, 2, 3, 4"
	B := "1, 2, 3"

	result := Sublist(A, B)

	expected := "A is a superlist of B"

	if strings.Contains(result, expected) == false {
		t.Errorf("Resultado %s != %s", result, expected)
	}

}

func TestNot(t *testing.T) {

	A := "1, 2, 4"
	B := "1, 2, 3"

	result := Sublist(A, B)

	expected := "A is not a superlist of, sublist of or equal to B"

	if strings.Contains(result, expected) == false {
		t.Errorf("Resultado %s != %s", result, expected)
	}

}
