package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

func (n *Node) GetVal() int {
	if n == nil {
		return 0
	}
	return n.val
}
func (n *Node) GetLeft() *Node {
	if n == nil {
		return nil
	}
	return n.left
}
func (n *Node) GetRight() *Node {
	if n == nil {
		return nil
	}
	return n.right
}

func merge(a, b *Node) *Node {
	if a == nil && b == nil {
		return nil
	}

	return &Node{
		val:   a.GetVal() + b.GetVal(),
		left:  merge(a.GetLeft(), b.GetLeft()),
		right: merge(a.GetRight(), b.GetRight()),
	}
}

func main() {

	left := &Node{
		left: &Node{
			val: 1,
			left: &Node{
				val: 2,
			},
		},
		right: &Node{
			val: 1,
		},
		val: 1,
	}
	right := &Node{
		left: &Node{
			val: 1,
		},
		right: &Node{
			val: 1,
		},
		val: 1,
	}

	merged := merge(left, right)
	fmt.Println(*merged)
}
