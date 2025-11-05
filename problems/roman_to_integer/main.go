package main

import "fmt"

func main() {

	s := "III"
	fmt.Println("res:", romanToInt(s))
}

func romanToInt(s string) int {

	nums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0

	for i := 0; i < len(s); i++ {

		fmt.Println("index", i)

		if i+1 < len(s) && nums[s[i]] < nums[s[i+1]] {
			res -= nums[s[i]]
		} else {
			res += nums[s[i]]
		}

		fmt.Printf("res: %d\n", res)

	}

	return res
}
