package main

import (
	"sort"
)

/*
Given an array of size n, find the majority element. The majority element is the
element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist
in the array.

Example 1:

Input: [3,2,3] Output: 3

Example 2:

Input: [2,2,1,1,1,2,2] Output: 2
*/

func MajorityElement(nums []int) int {
	// sort array
	// keep longest chain
	sort.Ints(nums)

	count := 0
	var k int

	for i := 0; i < len(nums); i++ {
		if nums[i] != k {
			count = 0
			k = nums[i]
		} else {
			count++
		}
		if count >= len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}

func MajorityElementMoores(nums []int) int {
	// Use Boyer-Moore majority vote algorithm
	// There can only be one.
	var m int
	i := 0

	for _, x := range nums {
		if i == 0 {
			m = x
			i++
		} else if m == x {
			i++
		} else {
			i--
		}
	}

	return m
}

/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:

Input: [3,2,3]
Output: [3]

Example 2:

Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
*/
func MajorityElements(nums []int) []int {
	// Use Boyer-Moore majority vote algorithm
	// There can only be one.
	var m1 int
	var m2 int
	i1 := 0
	i2 := 0

	for _, x := range nums {
		if m1 == x {
			i1++
		} else if m2 == x {
			i2++
		} else if i1 == 0 {
			m1 = x
			i1++
		} else if i2 == 0 {
			m2 = x
			i2++
		} else {
			i1--
			i2--
		}
	}

	// got two most common elements
	// now need to verify > 1/3
	cnt1 := 0
	cnt2 := 0
	for _, x := range nums {
		if x == m1 {
			cnt1++
		} else if x == m2 {
			cnt2++
		}
	}
	result := []int{}

	if cnt1 > len(nums)/3 {
		result = append(result, m1)
	}
	if cnt2 > len(nums)/3 {
		result = append(result, m2)
	}

	return result
}

/*
421. Maximum XOR of Two Numbers in an Array Medium

Given a non-empty array of numbers, a0, a1, a2, … , an-1, where 0 ≤ ai < 231.

Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.

Could you do this in O(n) runtime?

Example:

Input: [3, 10, 5, 25, 2, 8]

Output: 28

Explanation: The maximum result is 5 ^ 25 = 28.
*/
func findMaximumXORNaive(nums []int) int {
	var max int

	for i, a := range nums {
		for j := i + 1; j < len(nums); j++ {
			xor := a ^ nums[j]
			if xor > max {
				max = xor
			}
		}
	}

	return max
}
