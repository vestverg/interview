package main

import "fmt"

func solve(in map[int][]int) int {

	var res [][]int
	visited := make(map[int]bool)

	var queue []int
	for k := range in {

		if _, ok := visited[k]; ok {
			continue
		}
		var friendsGroup []int
		queue = append(queue, k)
		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			if visited[next] {
				continue
			}

			friendsGroup = append(friendsGroup, next)
			visited[next] = true
			if friends, ok := in[next]; ok {
				queue = append(queue, friends...)
			}
			delete(in, next)
		}
		res = append(res, friendsGroup)
	}

	fmt.Println(res)
	return len(res)
}

func main() {
	fmt.Println(solve(map[int][]int{
		0: {1, 2},
		1: {0, 5},
		2: {0},
		3: {6},
		4: {},
		5: {1},
		6: {3},
	}))
}
