/**
 *  最长回文子串
 *  给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string = "babad"
	substr := longestPalindrome(s)
	fmt.Printf("最长回文子串:%s", substr)
}

func longestPalindrome(s string) string {
	// 状态：f[i][j] 表示i到j是一个回文字符串
	// 推导：f[i][j] a[i]==a[j]&&f[i+1][j-1]==true
	// 初始化：f[j][j] = true
	// 结果：max(i-j,f[i][j]==true)
	if len(s) == 0 {
		return ""
	}
	res := make([]string, 0)
	res = append(res, s[0:1])
	f := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		f[i] = make([]bool, len(s))
	}
	f[0][0] = true
	for j := 1; j < len(s); j++ {
		f[j][j] = true
		for i := 0; i < j; i++ {
			// 推导过程
			if s[j] == s[i] && (i+1 >= j-1 || f[i+1][j-1]) {
				f[i][j] = true
				res = append(res, s[i:j+1])
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})
	return res[len(res)-1]
}
