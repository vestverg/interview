package main

import "fmt"

func solve(k int, arr []int) bool {

	freq := make(map[int]int)

	for _, val := range arr {
		freq[val] += 1
	}

	for _, v1 := range arr {
		freq[v1] -= 1
		t2 := k - v1
		for _, v2 := range arr {
			if freq[v2] > 0 {
				freq[v2] -= 1
				t3 := t2 - v2
				if freq[t3] > 0 {
					return true
				}
				freq[v2] += 1
			}
		}
		freq[v1] += 1

	}

	return false
}

func main() {
	fmt.Println(solve(49, []int{20, 303, 3, 4, 25}))
}
