/**
 * 最长上升子序列
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 * 注意: 因为题目中没有要求连续连续，LIS可能是连续的，也可能是非连续的
 */
package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("最长上升子序列的长度为%d\n", lengthOfLIS(nums))
}

/**
 * 动态规划 O(n2)
 * dp[i] ：表示以nums[i]结尾的最长上升子序列的长度
 * 状态转移方程为:dp[i] = max(dp[j]+1，dp[k]+1，dp[p]+1，.....)
 */
func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { //进行比较，因为dp[i]前面比他小的元素，不一定只有一个，找到最大的dp[j]
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

/**
 * 优化： O(nlogn)
 */
func lengthOfLIS1(nums []int) int {
	//todo
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
