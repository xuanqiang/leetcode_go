/**
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray2(nums)
	fmt.Printf("连续子数组最大为:%d\n", result)
}

/**
 * 动态规划
 * dp[i]: 表示以nums[i]结尾的连续子数组的最大和
 * 状态转移方程: dp[i]=max(nums[i], dp[i−1]+nums[i])
 * dp[0] = nums[0]
 * 最后寻找的是：max(dp[0], dp[1], ..., d[i-1], dp[i])
 */
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	//设置初始化值
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	//定义最小值
	result := math.MinInt32
	for _, key := range dp {
		result = max(result, key)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * 分治方法(典型)
 * 1、定义基本情况。
 * 2、将问题分解为子问题并递归地解决它们。
 * 3、合并子问题的解以获得原始问题的解。
 */
func maxSubArray1(nums []int) int {
	return 0
}

/**
 * 贪心算法
 * 使用单个数组作为输入来查找最大（或最小）元素（或总和）的问题，
 * 贪心算法是可以在线性时间解决的方法之一。
 * 每一步都选择最佳方案，到最后就是全局最优的方案。
 *
 * 该算法通用且简单：遍历数组并在每个步骤中更新：
 * 当前元素
 * 当前元素位置的最大和
 * 迄今为止的最大和
 */
func maxSubArray2(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}
