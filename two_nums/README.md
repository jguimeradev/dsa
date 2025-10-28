# TWO_NUMS

## Problem:

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.
 

## Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

## Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

## Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]


## Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


## Solution 1
```go
func twoSum(nums []int, target int) []int {
	for k, _ := range nums {
		for i := k + 1; i < len(nums); i++ {
			if nums[i] == target - nums[k] {
				return []int{k,i}
			}
		}
	}

	return []int{}
}
```
Time:  O(n2) -> Quadratic: The algorithm’s execution time grows proportionally to the square of the input size. Ex: 1, 2, 4, 8, 16, 32...
Space: O(1) -> constant space - same amount of memory regardless the input size -



