package main

import "fmt"

func countValidSelections(nums []int) int {

	const (
		R string = "right"
		L string = "left"
	)

	var index int = 3
	var direction string = R
	count := 0
	zeros := make(map[int]int)

	for index >= 0 && index <= len(nums)-1 {

		if len(zeros) == len(nums) {
			break
		}

		if nums[index] == 0 {

			if _, ok := zeros[index]; !ok {
				zeros[index] = 0
				// fmt.Println("zeros", zeros)
			}

			if direction == R {
				index++
				// fmt.Println("index move to right", index, nums[index])
			} else if direction == L {
				index--
				// fmt.Println("index move to left", index, nums[index])
			}

		} else if nums[index] > 0 {

			nums[index]--

			if _, ok := zeros[index]; !ok && nums[index] == 0 {
				zeros[index] = 0
				// fmt.Println("zeros", zeros)
			}

			// fmt.Println("decrement number", index, nums[index])
			if direction == R {
				direction = L
				index--
				// fmt.Println("index move to left", index, nums[index])
			} else {
				direction = R
				index++
				// fmt.Println("index move to right", index, nums[index])
			}
		}

		fmt.Println("nums", nums)
	}

	if index >= len(nums) || index < 0 && direction == L {
		count++
	}

	return 0
}

func main() {

	nums := []int{1, 0, 2, 0, 3}
	if countValidSelections(nums) == 0 {
		fmt.Println("successful")
	} else {
		fmt.Println("failure")
	}

}
