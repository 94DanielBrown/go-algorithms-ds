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

func Test_reversesString(t *testing.T) {
	result := reverse("abcd")
	if result != "dcba" {
		t.Errorf("Functon does not correctly reverse string")
	}
}
