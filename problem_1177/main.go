package main

import (
	"fmt"
	"math"
)

// k - number of times other elements appear
func solve(k int, arr ...int) int {
	var res int

	// iterate over all bits in number
	for i := 0; i < 32; i++ {
		count := 0
		for _, el := range arr {
			// count all bit in position i
			mask := 1 << i //
			bit := abs(el) & mask
			if bit != 0 { // if the bit at position i != 0
				count += 1
			}
		}
		// eliminate all bit that appear k times
		count = count % k

		res |= int(math.Pow(2.0, float64(i))) * count

		// put the bit in position count in i res
		//fmt.Println("count=", count, " ", strconv.FormatInt(int64(res), 2))
	}
	for _, i := range arr {
		if res == -1*i {
			res *= -1
		}
	}
	return res
}

func abs(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func main() {
	fmt.Println(solve(3, 5, 3, 3, 3, 6, 6, 6))
	fmt.Println(solve(3, 19, 13, 19, 19))
}
