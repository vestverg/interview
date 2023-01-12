package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func (s Set[T]) Add(val T) {
	s[val] = true
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s[val]
	return ok
}

type Transition struct {
	from        string
	to          string
	probability float64
}

//You are given a starting state start, a list of transition probabilities for a Markov chain,
//and a number of steps num_steps. Run the Markov chain starting from start for num_steps and compute the number of times we visited each state.
//
//For example, given the starting state a, number of steps 5000, and the following transition probabilities:
//
//[
//('a', 'a', 0.9),
//('a', 'b', 0.075),
//('a', 'c', 0.025),
//('b', 'a', 0.15),
//('b', 'b', 0.8),
//('b', 'c', 0.05),
//('c', 'a', 0.25),
//('c', 'b', 0.25),
//('c', 'c', 0.5)
//]
//One instance of running this Markov chain might produce { 'a': 3012, 'b': 1656, 'c': 332 }.

func solve(start string, count int, transitions []Transition) map[string]int {
	tr := make(map[string][]Transition)

	for _, transition := range transitions {
		tr[transition.from] = append(tr[transition.from], transition)
	}

	for _, v := range tr {
		sort.Slice(v, func(i, j int) bool {
			return v[i].probability < v[j].probability
		})
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s1)

	result := make(map[string]int)
	state := start
	for i := 0; i < count; i++ {
		next := getTransition(tr, state, random)
		result[next] += 1
	}
	return result
}

func getTransition(tr map[string][]Transition, state string, random *rand.Rand) string {
	r := random.Float64()

	neighbours := tr[state]
	curr := 0.0
	for _, transition := range neighbours {
		curr += transition.probability
		if curr >= r {
			return transition.to
		}

	}
	return ""
}

func main() {
	res := solve("a", 5000, []Transition{
		{"a", "a", 0.9},
		{"a", "b", 0.075},
		{"a", "c", 0.025},
		{"b", "a", 0.15},
		{"b", "b", 0.8},
		{"b", "c", 0.05},
		{"c", "a", 0.25},
		{"c", "b", 0.25},
		{"c", "c", 0.5},
	})
	fmt.Println(res)
}
