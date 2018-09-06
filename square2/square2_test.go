package main

import (
	"reflect"
	"testing"
)

func TestMakeList(t *testing.T) {

	result := MakeList(5)

	expected := []int{1, 2, 3, 4, 5}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Resultado %d é diferente do %d", result, expected)
	}

}
func TestSquareList(t *testing.T) {

	result := SquareList([]int{1, 2, 3, 4, 5})

	expected := []int{1, 4, 9, 16, 25}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Resultado %d é diferente do %d", result, expected)
	}
}

func TestSumList(t *testing.T) {

	result := SumList([]int{1, 2, 3, 4, 5})

	expected := 15

	if result != expected {
		t.Errorf("Resultado %d é diferente de %d", result, expected)
	}
}

func TestProcess(t *testing.T) {

	result := Process(10)

	expected := 2640

	if result != expected {
		t.Errorf("Resultado %d é diferente de %d", result, expected)
	}

}
