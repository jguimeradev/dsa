# TWO_NUMS

## Problem:

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
 

## Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.

## Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

## Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

## Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].


## Solution 1

```go
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
				res += I
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
```
