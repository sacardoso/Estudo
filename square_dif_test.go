package main

import (
	"testing"
)

var num, num1, num2 int
var result int
var expected int

func TestSquare(t *testing.T) {

	result := num * num

	if num == 1 {
		expected = 1
	}
	if num == 2 {
		expected = 4
	}
	if num == 3 {
		expected = 9
	}

	if result != expected {
		t.Errorf("%d != %d", result, expected)
	}
}

func TestSumNum(t *testing.T) {

	result := num1 + num2

	if num1 == 1 && num2 == 2 {
		expected = 3
	}

	if num1 == 2 && num2 == 2 {
		expected = 4
	}

	if num1 == 1 && num2 == 3 {
		expected = 4
	}

	if result != expected {
		t.Errorf("%d != %d", result, expected)
	}
}

func TestSqrNum(t *testing.T) {
	sumNum := num1 + num2
	result := sumNum * sumNum

	if num1 == 1 && num2 == 1 {
		expected = 4
	}
	if num1 == 1 && num2 == 2 {
		expected = 9
	}
	if num1 == 2 && num2 == 2 {
		expected = 16
	}

	if result != expected {
		t.Errorf("%d != %d", result, expected)
	}
}

func TestDif(t *testing.T) {

	sqr1 := num1 * num1
	sqr2 := num2 * num2
	sumNum := num1 + num2
	sqrSum := sumNum * sumNum
	sumSqr := sqr1 + sqr2
	result := sqrSum - sumSqr

	if sqrSum == 4 && sumSqr == 2 {
		expected = 2
	}

	if result != expected {
		t.Errorf("%d != %d", result, expected)
	}
}
