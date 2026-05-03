# 26. Remove Duplicates from Sorted Array

**Difficulty:** Easy

---

## Problem

Given an integer array `nums` sorted in non-decreasing order, remove the duplicates **in-place** such that each unique element appears only once. The relative order of the elements should be kept the same.

Return `k` — the number of unique elements. The first `k` elements of `nums` must contain the unique values in sorted order.

---

## Examples

**Example 1:**
```
Input:  nums = [1,1,2]
Output: 2, nums = [1,2,_]
```

**Example 2:**
```
Input:  nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
```

---

## Constraints

- `1 <= nums.length <= 3 * 10⁴`
- `-100 <= nums[i] <= 100`
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

---

### Mục tiêu đơn giản

Cho mảng số nguyên **đã sắp xếp tăng dần**, hãy **xóa các phần tử trùng lặp** ngay trên mảng gốc, sao cho mỗi giá trị chỉ xuất hiện **đúng 1 lần**.

---

### Điều quan trọng cần hiểu

**1. Không được tạo mảng mới**
> Phải xử lý **in-place** — tức là thao tác trực tiếp trên mảng đầu vào `nums`.

**2. Không cần xóa thật sự**
> Đề không yêu cầu mảng phải ngắn lại. Chỉ cần đảm bảo `k` phần tử đầu tiên chứa các giá trị unique theo thứ tự. Phần còn lại **không quan tâm**.

```
[0, 0, 1, 1, 2]  →  [0, 1, 2, _, _]
                             ↑
                     Bỏ qua phần này, đề không kiểm tra
```

## Cách giải bài 26 — Two Pointers

---

### Ý tưởng

Dùng **2 con trỏ** `i` và `j`:
- `i` → đánh dấu vị trí cuối vùng unique
- `j` → duyệt tìm phần tử mới

Vì mảng **đã sắp xếp**, phần tử trùng luôn nằm cạnh nhau → chỉ cần so sánh `nums[j]` với `nums[i]`.

---

### Các bước thực hiện

```
Bước 1: i = 0, j = 1

Bước 2: Với mỗi j từ 1 → n-1:
    Nếu nums[j] != nums[i]:
        i++
        nums[i] = nums[j]   ← ghi phần tử mới vào vùng unique

Bước 3: return i + 1
```

---

### Minh họa từng bước

```
nums = [0, 0, 1, 1, 1, 2]

       i=0
j=1 → nums[1]=0 == nums[0]=0  → bỏ qua
j=2 → nums[2]=1 != nums[0]=0  → i=1, nums[1]=1   [0,1,1,1,1,2]
j=3 → nums[3]=1 == nums[1]=1  → bỏ qua
j=4 → nums[4]=1 == nums[1]=1  → bỏ qua
j=5 → nums[5]=2 != nums[1]=1  → i=2, nums[2]=2   [0,1,2,1,1,2]

Kết quả: k = i+1 = 3
         nums = [0, 1, 2, _, _, _]
```

---

### Code

**Golang**
```go
func removeDuplicates(nums []int) int {
if len(nums) == 0 {
return 0
}

i := 0
for j :=1; j < len(nums); j++ {
if nums[j] != nums[i] {
i++
nums[i] = nums[j]
}
}

return i + 1
}
```

### Độ phức tạp

| | Độ phức tạp |
|---|---|
| **Thời gian** | O(n) — duyệt 1 lần |
| **Bộ nhớ** | O(1) — không dùng mảng phụ |