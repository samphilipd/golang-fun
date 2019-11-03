package main

import (
	"math"
)

/*
5. Longest Palindromic Substring Medium

Given a string s, find the longest palindromic substring in s. You may
assume that the maximum length of s is 1000.

Example 1:

Input: "babad" Output: "bab" Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd" Output: "bb"
*/

// ---

// 1221. Split a String in Balanced Strings
// Easy

// Balanced strings are those who have equal quantity of 'L' and 'R' characters.

// Given a balanced string s split it in the maximum amount of balanced strings.

// Return the maximum amount of splitted balanced strings.

// Example 1:

// Input: s = "RLRRLLRLRL"
// Output: 4
// Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

// Example 2:

// Input: s = "RLLLLRRRLR"
// Output: 3
// Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

// Example 3:

// Input: s = "LLLLRRRR"
// Output: 1
// Explanation: s can be split into "LLLLRRRR".

// Constraints:

//     1 <= s.length <= 1000
//     s[i] = 'L' or 'R'
func balancedStringSplit(s string) int {
	ticker := 0
	count := 0

	for _, c := range s {
		if c == 'R' {
			ticker++
		} else {
			ticker--
		}
		if ticker == 0 {
			count++
		}
	}
	return count
}

/*
The function first discards as many whitespace characters as necessary
until the first non-whitespace character is found. Then, starting from
this character, takes an optional initial plus or minus sign followed by
as many numerical digits as possible, and interprets them as a numerical
value.

The string can contain additional characters after those that form the
integral number, which are ignored and have no effect on the behavior of
this function.

If the first sequence of non-whitespace characters in str is not a valid
integral number, or if no such sequence exists because either str is
empty or it contains only whitespace characters, no conversion is
performed.

If no valid conversion could be performed, a zero value is returned.

Note:

- Only the space character ' ' is considered as whitespace character.
- Assume we are dealing with an environment which could only store
  integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. If
  the numerical value is out of the range of representable values,
  INT_MAX (2^31 − 1) or INT_MIN (−2^31) is returned.
*/
func myAtoi(str string) int {
	str = discardLeadingWhitespace(str)
	if str == "" {
		return 0
	}
	isPositive := true

	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		isPositive = false
		str = str[1:]
	}

	var result int32
	var underflow bool
	for _, rn := range str {
		n := int32(rn) - '0'
		if n >= 0 && n <= 9 {
			result, underflow = mul32(10, result)
			if !underflow {
				if isPositive {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			result, underflow = add32(result, n)
			if !underflow {
				if isPositive {
					return math.MaxInt32
				}
				return math.MinInt32
			}
		} else {
			break
		}
	}

	if isPositive {
		return int(result)
	}
	return -int(result)
}

func mul32(a, b int32) (int32, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}

func add32(a, b int32) (int32, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

func discardLeadingWhitespace(str string) string {
	for i, rn := range str {
		if rn != ' ' {
			return str[i:]
		}
	}
	return ""
}
