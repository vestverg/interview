package main

import (
	"fmt"
	"math"
)

// ordered indexes
func solve(tree []int, target int) []int {
	if tree[0] == target {
		return []int{}
	}
	level := int(math.Log2(float64(target)))
	from := int(math.Pow(2, float64(level))) - 1
	to := int(math.Pow(2, float64(level+1))) - 1
	var res []int
	parent := (target - 2) / 2
	for i := from; i <= to && i < len(tree); i++ {
		if tree[i] != target && tree[i] > 0 {
			if parent != (tree[i]-2)/2 {
				res = append(res, tree[i])
			}
		}
	}
	return res
}

type Node struct {
	left  *Node
	right *Node
	idx   int
}

// random ordered indexes
func solve2(root *Node, target int) []int {
	if root == nil && root.idx == target {
		return []int{}
	}

	parent, level := getLevel(root, target, 0)

	if parent < 0 || level < 0 {
		return []int{}
	}
	var res []int
	current := 0
	var nodes []*Node
	nodes = append(nodes, root)
	for len(nodes) != 0 && current <= level {
		var levelNodes []*Node

		for _, node := range nodes {
			if current != level {
				if node.left != nil {
					levelNodes = append(levelNodes, node.left)
				}
				if node.right != nil {
					levelNodes = append(levelNodes, node.right)
				}
			} else {
				if node.left != nil && node.idx != parent {
					res = append(res, node.left.idx)
				}
				if node.right != nil && node.right.idx != parent {
					res = append(res, node.right.idx)
				}
			}
			nodes = levelNodes
			current++
		}
	}
	return res
}

func getLevel(root *Node, target int, current int) (int, int) {

	if root == nil {
		return -1, -1
	}
	if root.left != nil && root.left.idx == target {
		return root.idx, current + 1
	}
	if root.right != nil && root.right.idx == target {
		return root.idx, current + 1
	}

	parent, left := getLevel(root.left, target, current+1)
	if left > 0 {
		return parent, left
	}
	parent, right := getLevel(root.right, target, current+1)

	if right > 0 {
		return parent, right
	}
	return -1, -1
}

func main() {
	fmt.Println(solve([]int{1, 2, 3, 4, 5, 6, 7}, 4))

	tree := &Node{
		left: &Node{
			left: &Node{
				left:  nil,
				right: nil,
				idx:   4,
			},
			right: &Node{
				left:  nil,
				right: nil,
				idx:   5,
			},
			idx: 2,
		},
		right: &Node{
			left: &Node{
				left:  nil,
				right: nil,
				idx:   7,
			},
			right: &Node{
				left:  nil,
				right: nil,
				idx:   6,
			},
			idx: 3,
		},
		idx: 1,
	}

	fmt.Println(solve2(tree, 5))
}
