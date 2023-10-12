package main

import (
	"reflect"
	"testing"
)

func Test_functionExists(t *testing.T) {
	if reflect.ValueOf(reverse).IsNil() {
		t.Errorf("Function doesn't exist")
	}
}

func Test_reversesInteger(t *testing.T) {
	reverseTests := []struct {
		input    int16
		expected int16
	}{
		{15, 51},
		{981, 189},
		{500, 5},
		{-15, -51},
		{-90, -9},
	}

	for _, e := range reverseTests {
		result := reverse(e.input)
		if result != e.expected {
			t.Errorf("Got %n but got %n", result, e.expected)
		}
	}
}
