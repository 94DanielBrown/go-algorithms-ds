package main

import (
	"reflect"
	"testing"
)

func Test_functionExists(t *testing.T) {
	if reflect.ValueOf(capitalize).IsNil() {
		t.Errorf("Function doesn't exist")
	}
}

func Test_chunksArray(t *testing.T) {
	capitalizeTests := []struct {
		string   string
		expected string
	}{
		{"a short sentence", "A Short Sentence"},
		{"a lazy fox", "A Lazy Fox"},
		{"look, it is working!", "Look, It Is Working!"},
	}

	for _, e := range capitalizeTests {
		result := capitalize(e.string)
		if result != e.expected {
			t.Errorf("Expected %s but got %s", e.expected, result)
		}
	}
}
