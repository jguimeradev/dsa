# Algorithms

## Binary Search




## Solution 1

```go
func binarySearch(nums []int, target int) int {

	low := 0
	up := len(nums) - 1

	for low <= up {
		mid := (low + up) / 2
		fmt.Println("mid", mid)
		if target == nums[mid] {
			return nums[mid]
		}
		if target < nums[mid] {
			up = mid - 1
		} else {
			low = mid + 1
		}

	}

	return 0

}
```

Time complexity: O(n).

## Solution 2


