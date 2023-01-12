package main

import (
	"fmt"
	"math"
)

func solve(tree string) int {

	if tree == "" {
		return 0
	}
	max := 0
	curr := 0
	for _, el := range tree {
		if el == '(' {
			curr++
			max = int(math.Max(float64(max), float64(curr)))
		}
		if el == ')' {
			curr--
		}
	}
	if curr != 0 {
		panic("invalid tree")
	}
	return max
}

func main() {
	fmt.Println(solve("((((00)0)0)0)"))
	fmt.Println(solve("(00)"))
	fmt.Println(solve("((00)(00))"))
	fmt.Println(solve("((00)(0(0(((00)0))(00)))))"))
}
