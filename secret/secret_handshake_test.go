package main

import (
	"reflect"
	"testing"
)

func TestHandshake(t *testing.T) {

	result := Secret(1)

	expected := []string{"wink"}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}

func TestHandshake2(t *testing.T) {

	result := Secret(3)

	expected := []string{"wink", "double blink"}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}

func TestHandshake16(t *testing.T) {

	result := Secret(16)

	expected := []string{}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}

func TestHandshake16plus(t *testing.T) {

	result := Secret(19)

	expected := []string{"double blink", "wink"}

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("Result %s != expected %s", result, expected)
	}
}
