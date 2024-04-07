package main

import "fmt"

func main() {
	node := Node{
		Data: 2,
	}

	node.add(3)

	node.add(4)

	fmt.Println(node.Children)

	node.remove(3)
	fmt.Println(node.Children)

	tree := Tree{}
	tree.Root = &node

}

type Node struct {
	Data     int
	Children []Node
}

func (n *Node) add(data int) {
	n.Children = append(n.Children, Node{Data: data})
}

func (n *Node) remove(data int) {
	for index, child := range n.Children {
		if child.Data == data {
			n.Children = append(n.Children[:index], n.Children[index+1:]...)
		}
	}
}

type Tree struct {
	Root *Node
}
