You are given an integer array nums.

Start by selecting a starting position curr such that nums[curr] == 0, and choose a movement direction of either left or right.

After that, you repeat the following process:

If curr is out of the range [0, n - 1], this process ends.
If nums[curr] == 0, move in the current direction by incrementing curr if you are moving right, or decrementing curr if you are moving left.
Else if nums[curr] > 0:
Decrement nums[curr] by 1.
Reverse your movement direction (left becomes right and vice versa).
Take a step in your new direction.
A selection of the initial position curr and movement direction is considered valid if every element in nums becomes 0 by the end of the process.

Return the number of possible valid selections.

## Iteration 1 

```go
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

```
## Output:
nums [1 0 0 0 1]
decrement number 4 0
index move to left 3 0
nums [1 0 0 0 0]
index move to left 2 0
nums [1 0 0 0 0]
index move to left 1 0
nums [1 0 0 0 0]
index move to left 0 1
nums [1 0 0 0 0]

## Example 2 (left)

NOT SOLVED YET
