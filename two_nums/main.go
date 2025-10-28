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
