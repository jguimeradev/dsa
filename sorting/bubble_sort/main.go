package main

import "fmt"

func main() {
	nums := [10]int{42, 15, 9, 77, 23, 50, 1, 35, 61, 88}
	bubbleSort(&nums)

}

func bubbleSort(nums *[10]int) {

	for i := 0; i < len(nums); i++ {
		order := false
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = swap(nums[j], nums[j+1])
				order = true //evita recorrer todo el array si ya está ordenado
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
