package main

import (
	"testing"
)

func Test_tree(t *testing.T) {
	node := Node{Data: 4}

	if node.Data != 4 {
		t.Errorf("Expected 4 but got %v", node.Data)
	}

	if len(node.Children) != 0 {
		t.Errorf("Expected 0 but got %v", node.Children)
	}

	node.add(1)
	node.add(7)

	if len(node.Children) != 2 {
		t.Errorf("Expected 2 but got %v", node.Children)
	}

	tree := Tree{Root: nil}
	tree.Root = &node
	if tree.Root == nil {
		t.Errorf("Expected memory address of node but got nil")
	}

}
