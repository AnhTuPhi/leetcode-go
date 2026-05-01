package main

import "leetcode-go/runner"

type Input struct {
	nums []int
	val  int
}

type Main struct {
}

func (m Main) Solve(input Input) []int {
	nums := input.nums
	val := input.val
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return nums
}

func (m Main) TestCases() []runner.TestCase[Input, []int] {
	return []runner.TestCase[Input, []int]{
		runner.Case("example 1", Input{[]int{3, 2, 2, 3}, 3}, []int{2, 2}),
		runner.Case("example 2", Input{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, []int{0, 1, 4, 0, 3}),
	}
}
func main() {
	main := Main{}
	runner.Run[Input, []int](main)
	//runner.Benchmark[Input, []int](main, 100, 1000)
}
