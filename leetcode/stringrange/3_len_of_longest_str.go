//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4233 👎 0
package stringrange

import (
	"fmt"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func LengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	bs := strings.Split(s, "")
	type Link struct {
		Val  string
		Next *Link
		Id   int
	}
	max := 0
	var first, last *Link
	for _, b := range bs {
		if first == nil {
			first = &Link{b, nil, 0}
			last = first
			continue
		}
		point := first
		for {
			fmt.Println("   ", point)
			if point.Val == b {
				if last.Id-first.Id > max {
					max = last.Id - first.Id
				}
				last.Next = &Link{b, nil, last.Id + 1}
				last = last.Next
				first = point.Next
				break
			}
			if point == last {
				last.Next = &Link{b, nil, last.Id + 1}
				last = last.Next
				break
			}
			point = point.Next
		}
		fmt.Println(last)
	}
	if last.Id-first.Id > max {
		max = last.Id - first.Id
	}
	return max + 1
}


//leetcode submit region end(Prohibit modification and deletion)
