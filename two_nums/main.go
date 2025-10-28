package main

import (
	"fmt"
)

func main() {

	nums := []int{11, 15, 11, 11, 11, 11, 11, 2, 7}
	target := 9

	res := twoSum(nums, target)

	fmt.Println("res", res)

}

func twoSum(nums []int, target int) []int {
	for k, _ := range nums {
		for i := k + 1; i < len(nums); i++ {
			if nums[i] == target-nums[k] {
				return []int{k, i}
			}
		}
	}

	return []int{}
}
