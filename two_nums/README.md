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
Time complexity:  O(n2) 
Quadratic: The algorithm’s execution time grows proportionally to the square of the input size. Ex: 1, 2, 4, 8, 16, 32...
Space complexity: O(1)
Constant space - same amount of memory regardless the input size -   


## Solution 2

```go
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for k, v := range nums {
		value := target - nums[k]
		if j, ok := m[value]; ok {
			return []int{k, j}
		}
		m[v] = k
	}

	return []int{}

}
```

Iterating and inserting elements into the hash table (map) and check if current element's value already exists in the hash table. If it exists, we have found a solution and return the indices immediately.

Time complexity: O(n).
We traverse the list containing n elements exactly twice. Since the hash table reduces the lookup time to O(1), the overall time complexity is O(n).

Space complexity: O(n).
The extra space required depends on the number of items stored in the hash table, which stores exactly n elements.


