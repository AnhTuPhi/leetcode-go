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

	i := 1
	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	i++
	return nums
}
