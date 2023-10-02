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
