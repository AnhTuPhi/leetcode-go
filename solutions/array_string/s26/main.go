package main

type Input struct {
	nums []int
}

type Main struct {
}

func (m Main) Solve(input Input) []int {
	nums := input.nums

	if len(nums) == 0 {
		return []int{}
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	i++
	return nums
}
