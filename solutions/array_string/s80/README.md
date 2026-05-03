# 80. Remove Duplicates from Sorted Array II

**Difficulty:** Medium

---

## Problem

Given an integer array `nums` sorted in non-decreasing order, remove some duplicates **in-place** such that each unique element appears **at most twice**. The relative order of the elements should be kept the same.

Return `k` — the number of elements after removing duplicates. The first `k` elements of `nums` must contain the final result.

> Do not allocate extra space. You must do this **in-place** with **O(1)** extra memory.

---

## Examples

**Example 1:**
```
Input:  nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
```

**Example 2:**
```
Input:  nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
```

---

## Constraints

- `1 <= nums.length <= 3 * 10⁴`
- `-10⁴ <= nums[i] <= 10⁴`
- `nums` is sorted in **non-decreasing** order

---

## Custom Judge

```java
int[] nums = [...];         // Input array
int[] expectedNums = [...]; // Expected answer with correct length

int k = removeDuplicates(nums);

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

## Cách giải

```go
func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }

	i := 1
	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}

    return i + 1
}
```