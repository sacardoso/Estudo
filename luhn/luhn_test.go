package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoubleNumber(t *testing.T) {

	result := DoubleNumber([]int{9, 8, 7, 6, 5})

	expected := []int{9, 8, 5, 6, 1}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Resultado %d é diferente do %d", result, expected)
	}
}

func TestSumDigits(t *testing.T) {

	numeros := []int{1, 2, 3, 4, 5}

	result := SumDigits(numeros)

	expected := 15

	if result != expected {
		t.Errorf("Resultado %d é diferente do %d", result, expected)
	}
}

func TestVerifyIfNumberIsValid(t *testing.T) {

	result := VerifyIfNumberIsValid(30)

	expected := "The number is valid!"

	if result != expected {
		fmt.Println("Resultado:", result, "é diferente de:", expected)
	}

}
