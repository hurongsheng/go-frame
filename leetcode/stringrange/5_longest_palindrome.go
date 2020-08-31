package stringrange

import "fmt"

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
// 示例 1：
//
// 输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//
//
// 示例 2：
//
// 输入: "cbbd"
//输出: "bb"
//
// Related Topics 字符串 动态规划
// 👍 2628 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func LongestPalindrome(s string) string {
	if s == "" {
		return s
	}
	for l := len(s) - 1; l >= 0; l-- {
		for start := 0; start < len(s); start++ {
			end := start + l
			if end >= len(s) {
				fmt.Println("break2")
				break
			}
			j := 0
			flag := true
			for {
				fmt.Printf("try %v(%v,%v) \n", s[start:end+1], j, l)
				if s[start+j] != s[end-j] {
					fmt.Printf("false \n")
					flag = false
					break
				}
				if start+j+1 > end-j-1 {
					fmt.Printf("break \n")
					break
				}
				j++
			}
			if flag {
				fmt.Printf("res %v \n", s[start:end+1])
				return s[start : end+1]
			}
		}
	}
	return s[0:1]
}

//leetcode submit region end(Prohibit modification and deletion)
