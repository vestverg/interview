package main

import (
	"fmt"
)

func solve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	f0 := 0
	f1 := 1
	for i := 2; i <= n; i++ {
		t := f0 + f1
		f0 = f1
		f1 = t
	}
	return f1
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("n=", i, " ", solve(i))
	}
}
