package main

import (
	"testing"
)

func Test_findsMaxChar(t *testing.T) {
	maxCharTests := []struct {
		input    string
		expected string
	}{
		{"abccccccccd", "c"},
		{"apple 1231111", "1"},
	}

	for _, e := range maxCharTests {
		result := maxChars(e.input)
		if result != e.expected {
			t.Errorf("Got %v but got %v", result, e.expected)
		}
	}
}
