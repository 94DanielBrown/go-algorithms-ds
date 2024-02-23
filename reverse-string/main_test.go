package main

import (
	"testing"
)

func Test_reversesString(t *testing.T) {
	result := reverse("abcd")
	if result != "dcba" {
		t.Errorf("Functon does not correctly reverse string")
	}
}
