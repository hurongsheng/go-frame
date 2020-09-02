//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 示例 1: 
//
// 输入: ["flower","flow","flight"]
//输出: "fl"
// 
//
// 示例 2: 
//
// 输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
// 
//
// 说明: 
//
// 所有输入只包含小写字母 a-z 。 
// Related Topics 字符串 
// 👍 1239 👎 0

package stringrange

//leetcode submit region begin(Prohibit modification and deletion)
func LongestCommonPrefix(strs []string) string {
	i := 0
	if len(strs) == 0 {
		return ""
	}
	commPre := ""
	for {
		pref := ""
		for _, s := range strs {
			if i > len(s)-1 {
				return commPre
			}
			if pref == "" {
				pref = s[i : i+1]
				continue
			}
			if pref != s[i:i+1] {
				return commPre
			}

		}
		commPre += pref
		i++
	}

}

//leetcode submit region end(Prohibit modification and deletion)
