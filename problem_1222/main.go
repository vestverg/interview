package main

import "fmt"

func solve(tree []int) (int, int) {

	freq := make(map[int]int)
	_ = traverse(tree, 0, freq)
	val := 0
	max := 0
	for el, fr := range freq {
		if fr > max {
			val = el
			max = fr
		}
	}
	return val, max
}

func traverse(tree []int, node int, freq map[int]int) int {
	val := tree[node]
	if (2*node+1) > len(tree)-1 || (2*node+2) > len(tree)-1 {
		freq[val] += 1
		return val
	}
	left := traverse(tree, 2*node+1, freq)
	right := traverse(tree, 2*node+2, freq)
	res := val + left + right
	freq[res] += 1
	return res
}

func main() {
	fmt.Println(solve([]int{5, 2, -5}))
	fmt.Println(solve([]int{5, 2, -5, 1, 1, 2, 3}))
}
