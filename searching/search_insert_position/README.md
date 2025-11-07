# SEARCH INSERT POSITION

## Problem:

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order

## Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2

## Example 2:
Input: nums = [1,3,5,6], target = 2
Output: 1

## Example 3:
Input: nums = [1,3,5,6], target = 7
Output: 4


## Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104

## Solution 1

```go
func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {

		mid := (high + low) / 2

        if (target == nums[mid]){
            return mid
        }

		if target < nums[0] {
			return 0
		} else if target > nums[high] {
			return high + 1
		}

		if target < nums[mid] {
			high = mid - 1
            if target > nums[high]{
                return high + 1
            }
		} else if target > nums[mid] {
			low = mid + 1
            if target < nums[low]{
                return low
            }
		}
	}

	return 0
}
```
Time complexity:  O(log n) 
 


## Solution 2

```go
func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {

		mid := (high + low) / 2

		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
            return mid
        }

	}
    return low
}
```
Explanation here:
[Link](https://www.youtube.com/watch?v=v4J_AWp-6EQ&t=359s)


