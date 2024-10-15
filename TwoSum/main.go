package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	var nums1 = []int{3, 2, 4}
	target1 := 6
	fmt.Println(twoSum(nums1, target1))
	var nums2 = []int{3, 3}
	target2 := 6
	fmt.Println(twoSum(nums2, target2))
}
