package main

import (
	"reflect"
	"testing"
)

func Test_functionExists(t *testing.T) {
	if reflect.ValueOf(anagrams).IsNil() {
		t.Errorf("Function doesn't exist")
	}
}

func Test_chunksArray(t *testing.T) {
	anagramTests := []struct {
		firstString  string
		secondString string
		expected     bool
	}{
		{"rail safety", "fairy tales", true},
		{"RAIL! SAFETY!", "fairy tales", true},
		{"Hi there", "Bye there", false},
	}

	for _, e := range anagramTests {
		result := anagrams(e.firstString, e.secondString)
		if result != e.expected {
			t.Errorf("Expected %v but got %v", e.expected, result)
		}
	}
}
