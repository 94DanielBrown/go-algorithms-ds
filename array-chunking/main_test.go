package main

import (
	"reflect"
	"testing"
)

func Test_functionExists(t *testing.T) {
	if reflect.ValueOf(chunk).IsNil() {
		t.Errorf("Function doesn't exist")
	}
}

func Test_chunksArray(t *testing.T) {
	chunkTests := []struct {
		array    []int8
		size     int8
		expected [][]int8
	}{
		{[]int8{1, 2, 3, 4, 5}, 2, [][]int8{{1, 2}, {3, 4}, {5}}},
	}

	for _, e := range chunkTests {
		result := chunk(e.array, e.size)
		if !reflect.DeepEqual(result, e.size) {
			t.Errorf("Got %n but got %n", result, e.expected)
		}
	}
}
