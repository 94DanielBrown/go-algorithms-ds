package main

import (
	"io"
	"os"
	"testing"
)

func Test_printsSteps(t *testing.T) {
	stepsTests := []struct {
		input    int
		expected string
	}{
		{1, "#\n"},
		//{2, "# \n##\n"},
		//{3, "# \n## \n###\n"},
		//{4, "# \n## \n### \n####\n"},
	}

	for _, e := range stepsTests {
		oldOut := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		printSteps(e.input)

		_ = w.Close()
		os.Stdout = oldOut
		out, _ := io.ReadAll(r)
		result := string(out)

		if result != e.expected {
			t.Errorf("Expected %s but got %s", e.expected, result)
		}
	}
}
