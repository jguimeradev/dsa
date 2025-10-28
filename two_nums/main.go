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

	sum := 0

	aux := []int{}

	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			sum = v + nums[i]
			if sum == target {
				aux = append(aux, k, i)
				return aux
			}
		}
	}

	return aux
}
