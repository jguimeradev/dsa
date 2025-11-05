# Bubble Sort Algorithm

## My solution

```go
func bubbleSort(nums [10]int) {

	for i := 0; i < len(nums); i++ {
		order := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = swap(nums[j], nums[j+1])
				order = true
			}
		}
		if !order {
			break
		}
	}

	fmt.Println(nums)

}

func swap(x int, y int) (int, int) {
	return y, x
}

```

## Idiomatic Go version:


```go
func bubbleSort(nums []int) {
    n := len(nums)
    for i := 0; i < n; i++ {
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
                swapped = true
            }
        }
        if !swapped {
            break // array is already sorted
        }
    }
}
```

