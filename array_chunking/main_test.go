package main

import (
	"reflect"
	"testing"
)

func Test_chunksArray(t *testing.T) {
	chunkTests := []struct {
		array    []int
		size     int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
	}

	for _, e := range chunkTests {
		result := chunk(e.array, e.size)
		if !reflect.DeepEqual(result, e.size) {
			t.Errorf("Got %n but got %n", result, e.expected)
		}
	}
}
