# 88. Merge Sorted Array

**Level:** Easy

## Problem

You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

Merge `nums1` and `nums2` into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array `nums1`. To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to `0` and should be ignored. `nums2` has a length of `n`.

---

## Examples

**Example 1:**
```
Input:  nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
```
> The arrays we are merging are `[1,2,3]` and `[2,5,6]`.  
> The result of the merge is `[1,2,2,3,5,6]` with the underlined elements coming from `nums1`.

**Example 2:**
```
Input:  nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
```
> The arrays we are merging are `[1]` and `[]`.  
> The result of the merge is `[1]`.

**Example 3:**
```
Input:  nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
```
> The arrays we are merging are `[]` and `[1]`.  
> The result of the merge is `[1]`.  
> Note that because `m = 0`, there are no elements in `nums1`. The `0` is only there to ensure the merge result can fit in `nums1`.

---

## Constraints

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`

---

## Follow Up

> Can you come up with an algorithm that runs in **O(m + n)** time?

---

## Solution

Bài toán này có thể được giải bằng nhiều cách, song mỗi cách đều có performance khác nhau về độ phức tạp.

**Cách tối ưu nhất:** Three Pointer — kỹ thuật dùng 3 con trỏ để duyệt mảng từ cuối.

| | Độ phức tạp |
|---|---|
| **Time** | O(m + n) |
| **Space** | O(1) — không cần mảng phụ |

**Ưu điểm:** Tận dụng việc 2 mảng ban đầu đã được sort tăng dần, không tốn thêm bộ nhớ.

### Ý tưởng

Đặt 3 con trỏ:
- `i = m - 1` → trỏ vào phần tử hợp lệ cuối cùng của `nums1`
- `j = n - 1` → trỏ vào phần tử cuối cùng của `nums2`
- `k = m + n - 1` → vị trí ghi (cuối `nums1`)

So sánh `nums1[i]` và `nums2[j]`, phần tử nào **lớn hơn** thì ghi vào `nums1[k]`, sau đó lùi con trỏ tương ứng.

### Code (Java)

```java
class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int i = m - 1;
        int j = n - 1;
        int k = m + n - 1;

        while (i >= 0 && j >= 0) {
            if (nums1[i] > nums2[j]) {
                nums1[k--] = nums1[i--];
            } else {
                nums1[k--] = nums2[j--];
            }
        }

        while (j >= 0) {
            nums1[k--] = nums2[j--];
        }
    }
}
```

### Minh họa

![Three Pointer Visualization](https://assets.leetcode.com/users/images/6e9bad1b-c9f7-44e1-86f5-eaa834a74632_1673465370.762181.jpeg)

### Từng bước với Example 1
```go
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
```

Minh hoạ
```
nums1 = [1, 2, 3, 0, 0, 0]   m = 3
nums2 = [2, 5, 6]             n = 3

i=2, j=2, k=5
Bước 1: nums1[2]=3 vs nums2[2]=6 → ghi 6 → [1,2,3,0,0,6]  | i=2, j=1, k=4
Bước 2: nums1[2]=3 vs nums2[1]=5 → ghi 5 → [1,2,3,0,5,6]  | i=2, j=0, k=3
Bước 3: nums1[2]=3 vs nums2[0]=2 → ghi 3 → [1,2,3,3,5,6]  | i=1, j=0, k=2
Bước 4: nums1[1]=2 vs nums2[0]=2 → ghi 2 → [1,2,2,3,5,6]  | i=1, j=-1, k=1

j < 0 → thoát vòng lặp
Kết quả: [1,2,2,3,5,6] ✅
```

![img.png](img.png)