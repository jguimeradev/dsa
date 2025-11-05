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

	nums := []int{46, 53, 45, 49, 53, 45, 47, 44, 46, 37, 50, 39, 53, 52, 48, 42, 40, 48, 43, 41, 51, 45, 41, 43, 49, 44, 45, 45, 43, 51, 51, 46, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 35, 39, 40, 41, 34, 40, 40, 32, 27, 27, 41, 40, 39, 42, 32, 52, 32, 29, 35, 40, 32, 37, 28, 34, 42, 33, 38, 40, 40, 40, 38, 38, 30, 38, 38, 43, 36, 36, 38, 40}
	if countValidSelections(nums) == 0 {
		fmt.Println("successful")
	} else {
		fmt.Println("failure")
	}

}
