package main

import "fmt"

func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {

		mid := (high + low) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[0] {
			return 0
		} else if target > nums[high] {
			return high + 1
		}

		if target < nums[mid] {
			high = mid - 1
			if target > nums[high] {
				return high + 1
			}
		} else if target > nums[mid] {
			low = mid + 1
			if target < nums[low] {
				return low
			}
		}
	}

	return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
