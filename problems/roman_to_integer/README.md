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

This solution runs in linear time (O(n)) but is inefficient in structure and style. However, it allowed me to practice `switch` statements and observe that in Go, `i++` and `i--` are allowed only as statements, not expressions. Therefore, `arr[i++]` fails at compile time, whereas `arr[i+1]` works.

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

## Solution 2

```go
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
```
## Explanation

It works because of how **Roman numerals encode subtraction**. The algorithm’s rule—**subtract if the current value is less than the next; otherwise add**—replicates the Roman numeral system’s positional logic.

---

### **1. Roman numeral principle**

Roman numerals are additive **unless** a smaller numeral precedes a larger one.
Examples:

* `VI` → 5 + 1 = 6 (additive order)
* `IV` → 5 − 1 = 4 (subtractive order)

The *direction* of the comparison determines addition or subtraction.

---

### **2. Algorithmic idea**

For a Roman string `s`, read from left to right:

* Let `v[i]` be the value of `s[i]`.
* If `v[i] < v[i+1]`, subtract `v[i]`.
* Else, add `v[i]`.

This works because only valid Roman numerals follow the subtractive rule in exactly those cases (I before V or X, X before L or C, C before D or M).

---

### **3. Example: MCMXCIV**

| Char | Value | Next | Rule | Operation | Result |
| ---- | ----- | ---- | ---- | --------- | ------ |
| M    | 1000  | C    | >    | add       | 1000   |
| C    | 100   | M    | <    | subtract  | 900    |
| M    | 1000  | X    | >    | add       | 1900   |
| X    | 10    | C    | <    | subtract  | 1890   |
| C    | 100   | I    | >    | add       | 1990   |
| I    | 1     | V    | <    | subtract  | 1989   |
| V    | 5     | —    | add  | 1994      |        |

---

### **4. Why it always works**

* Roman numerals are *monotonic* except at valid subtractive pairs.
* Each subtractive case uses exactly one smaller–before–larger pair.
* Thus, checking `current < next` correctly captures those pairs.

---

### **5. Complexity**

* **Time:** O(n) — one pass
* **Space:** O(1)

---

### **Summary**

The rule works because Roman numerals are **mostly additive** with **explicit subtractive exceptions**, and comparing adjacent values is enough to detect them. The algorithm is simple, linear, and exact for all valid Roman numerals.
