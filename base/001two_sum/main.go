/**
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 * 给定 nums = [2, 7, 11, 15], target = 9
 * 所以返回 [0, 1]
**/
package main

import "fmt"

func main() {
	var nums = []int{2, 5, 7, 15}
	var target int = 9
	fmt.Printf("打印结果为:%v", twoSum1(nums, target))
}

//暴力破解
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return nil
}

//运用哈希表,空间换时间
func twoSum1(nums []int, target int) []int {
	var result = []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
