package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_functionExists(t *testing.T) {
	if reflect.ValueOf(pyramid).IsNil() {
		t.Errorf("Function doesn't exist")
	}
}

func Test_pyramid(t *testing.T) {
	pyramidTests := []struct {
		n        int
		expected string
	}{
		{1, "#\n"},
		{2, "# \n###\n"},
		{3, "# \n### \n#####\n"},
	}

	for _, e := range pyramidTests {
		oldOut := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		pyramid(e.n)
		_ = w.Close()
		os.Stdout = oldOut
		out, _ := io.ReadAll(r)
		result := string(out)

		if result != e.expected {
			t.Errorf("Expected %s but got %s", e.expected, result)
		}

	}
}
