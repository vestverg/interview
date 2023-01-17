package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	from, to int
}

func solve(arr []Interval) []Interval {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].from < arr[j].from
	})
	st := []Interval{}
	for _, interval := range arr {
		if len(st) == 0 {
			st = append(st, interval)
			continue
		}
		last := st[len(st)-1]
		if overlap(last, interval) {
			st = st[:len(st)-1]
			merged := merge(last, interval)
			st = append(st, merged)
		} else {
			st = append(st, interval)
		}
	}
	return st
}

func overlap(i, j Interval) bool {
	return j.from >= i.from && j.from <= i.to
}

func merge(i, j Interval) Interval {
	to := i.to
	if i.to < j.to {
		to = j.to
	}
	return Interval{
		from: i.from,
		to:   to,
	}

}

func main() {
	intervals := []Interval{
		{1, 3},
		{5, 8},
		{4, 10},
		{20, 25},
	}
	result := solve(intervals)
	fmt.Println(result)
}
