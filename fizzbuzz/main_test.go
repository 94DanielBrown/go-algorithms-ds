package main

import (
	"io"
	"os"
	"testing"
)

func Test_fizzbuzz(t *testing.T) {
	fizzbuzzTests := []struct {
		n        int
		expected string
	}{
		{1, "1\n"},
		{5, "1\n2\nfizz\n4\nbuzz\n"},
		{15, "1\n2\nfizz\n4\nbuzz\nfizz\n7\n8\nfizz\nbuzz\n11\nfizz\n13\n14\nfizzbuzz\n"},
	}

	for _, e := range fizzbuzzTests {
		oldOut := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		fizzbuzz(e.n)
		_ = w.Close()
		os.Stdout = oldOut
		out, _ := io.ReadAll(r)
		result := string(out)

		if result != e.expected {
			t.Errorf("Expected %s but got %s", e.expected, result)
		}

	}
}
