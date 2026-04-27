package main

import (
	"leetcode-go/runner"
)

type Input struct {
	Nums1 []int
	m     int
	Nums2 []int
	n     int
}

type Main struct {
}

func (main Main) Solve(input Input) []int {
	Nums1 := input.Nums1
	m := input.m
	Nums2 := input.Nums2
	n := input.n

	i := m - 1
	j := n - 1
	k := len(Nums1) - 1

	for i >= 0 && j >= 0 {
		if Nums1[i] > Nums2[j] {
			Nums1[k] = Nums1[i]
			i--
		} else {
			Nums1[k] = Nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		Nums1[k] = Nums2[j]
		k--
		j--
	}

	return Nums1
}

func (m Main) TestCases() []runner.TestCase[Input, []int] {
	return []runner.TestCase[Input, []int]{
		runner.Case("example 1 - basic merge", Input{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}),
		runner.Case("example 2 - nums2 all smaller", Input{[]int{1}, 1, []int{}, 0}, []int{1}),
		runner.Case("example 3 - nums1 empty", Input{[]int{0}, 0, []int{1}, 1}, []int{1}),
	}
}

func main() {
	main := Main{}
	runner.Run[Input, []int](main)
	//runner.Benchmark[Input, []int](main, 100, 1000)
}
