/* Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target. You may assume that
each input would have exactly one solution, and you may not use the same
element twice. You can return the answer in any order. */

package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}
