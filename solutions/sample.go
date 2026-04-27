package main

import (
	"leetcode-go/runner"
	"sort"
)

// Input struct — gom nhiều param vào 1 type
type Input struct {
	Nums   []int
	Target int
}

type TwoSum struct{}

func (t TwoSum) Solve(input Input) []int {
	seen := make(map[int]int)
	for i, num := range input.Nums {
		complement := input.Target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func (t TwoSum) TestCases() []runner.TestCase[Input, []int] {
	return []runner.TestCase[Input, []int]{
		runner.Case("basic", Input{[]int{2, 7, 11, 15}, 9}, []int{0, 1}),
		runner.Case("middle", Input{[]int{3, 2, 4}, 6}, []int{1, 2}),
		runner.Case("same", Input{[]int{3, 3}, 6}, []int{0, 1}),
	}
}

// Output []int không đảm bảo thứ tự → sort trước khi so
func (t TwoSum) AssertEqual(expected, actual []int) bool {
	if len(expected) != len(actual) {
		return false
	}
	e, a := append([]int{}, expected...), append([]int{}, actual...)
	sort.Ints(e)
	sort.Ints(a)
	for i := range e {
		if e[i] != a[i] {
			return false
		}
	}
	return true
}

func main() {
	sol := TwoSum{}
	runner.Run[Input, []int](sol)
	runner.Benchmark[Input, []int](sol, 100, 1000)
}
