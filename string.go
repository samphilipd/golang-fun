package main

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
