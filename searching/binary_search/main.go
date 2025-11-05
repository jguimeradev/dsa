package main

import "fmt"

func main() {

	nums := []int{1, 3, 5, 8, 12, 14, 17, 19, 22, 25, 27, 30, 33, 37, 40}
	target := 40
	res := binarySearch(nums, target)
	fmt.Printf("res: %d, target: %d", res, target)

}

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
