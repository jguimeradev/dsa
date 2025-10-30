package main

import "fmt"

var (
	I int = 1
	V int = 5
	X int = 10
	L int = 50
	C int = 100
	D int = 500
	M int = 1000
)

func main() {

	s := "MMMCC"
	fmt.Println("res:", romanToInt(s))
}

func romanToInt(s string) int {

	var res int = 0

	i := 0
	for i < len(s) {
		switch s[i] {
		case 'I':
			if i < len(s)-1 {
				if s[i+1] == 'V' {
					res += 4
					i++
				} else if s[i+1] == 'X' {
					res += 9
					i++
				} else {
					res += I
				}
			} else {
				res += I
			}
		case 'V':
			res += V
		case 'X':
			if i < len(s)-1 {
				if s[i+1] == 'L' {
					res += 40
					i++
				} else if s[i+1] == 'C' {
					res += 90
					i++
				} else {
					res += X
				}
			} else {
				res += X
			}
		case 'L':
			res += L
		case 'C':
			if i < len(s)-1 {
				if s[i+1] == 'D' {
					res += 400
					i++
				} else if s[i+1] == 'M' {
					res += 900
					i++
				} else {
					res += C
				}
			} else {
				res += C
			}
		case 'D':
			res += D
		case 'M':
			res += M
		}
		i++
	}

	return res
}
