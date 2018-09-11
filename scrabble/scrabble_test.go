package main

import "testing"

func TestScrabble(t *testing.T) {

	result := ScrabbleScore("cabbage")

	expected := 14

	if result != expected {
		t.Errorf("Errado!")
	}
}

func TestScrabble2(t *testing.T) {

	result := ScrabbleScore("CaBbagE")

	expected := 14

	if result != expected {
		t.Errorf("Resultado: %d != Esperado: %d", result, expected)
	}
}

func TestTwoWords(t *testing.T) {

	result := ScrabbleScore("cabbage cabbage")

	expected := 28

	if result != expected {
		t.Errorf("Resultado: %d != Esperado: %d", result, expected)
	}

}

func TestZero(t *testing.T) {

	result := ScrabbleScore(" ")

	expected := 0

	if result != expected {
		t.Errorf("Resultado: %d != Esperado: %d", result, expected)
	}

}
