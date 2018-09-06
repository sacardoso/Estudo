package main

import (
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
