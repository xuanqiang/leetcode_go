/**
 * 70
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 */
package main

import "fmt"

func main() {
	var n int = 2
	fmt.Printf("爬的楼梯数为:%d\n", climbStairs(n))
}

// 令 dp[n] 表示能到达第 n 阶的方法总数
// 状态转移方程: dp[n]=dp[n-1]+dp[n-2]
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
