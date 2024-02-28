package main

import "testing"

func Test_linkedList(t *testing.T) {
	ll := LinkedList{}

	// Test inserting values
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	// Test finding values
	tests := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false}, // This value was not inserted, so should not be found
	}

	for _, test := range tests {
		result := ll.Find(test.value)
		if (result != nil) != test.expected {
			t.Errorf("Find(%d) = %v; want %v", test.value, result != nil, test.expected)
		}
	}
}
