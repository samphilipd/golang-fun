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
	sort.Ints(nums)

	count := 0
	var k int
	els := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] != k {
			count = 0
			k = nums[i]
		} else {
			count++
		}
		if count >= len(nums)/3 {
			els = append(els, nums[i])
		}
	}
	return els
}
