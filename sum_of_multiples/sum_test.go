package main

import (
	"reflect"
	"testing"
)

func TestMultiples(t *testing.T) {

	result := Multiples(2, 7, 10)

	expected := []int{2, 4, 6, 8, 7}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %d != Expected %d", result, expected)
	}
}

func TestMultiplesWithIntersection(t *testing.T) {

	result := Multiples(2, 3, 10)

	expected := []int{2, 4, 6, 8, 3, 0, 9}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %d != Expected %d", result, expected)
	}
}

func TestSumEqualNumbers(t *testing.T) {

	result := SumOfMultiples(2, 2, 2)

	expected := 0

	if result != expected {
		t.Errorf("Result %d != Expected %d", result, expected)
	}
}

func TestSum(t *testing.T) {

	result := SumOfMultiples(2, 3, 10)

	expected := 32

	if result != expected {
		t.Errorf("Result %d != Expected %d", result, expected)
	}
}
