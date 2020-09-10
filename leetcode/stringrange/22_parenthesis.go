//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。 
//
// 
//
// 示例： 
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//       "(())(())"
//     ]
// 
// Related Topics 字符串 回溯算法 
// 👍 1288 👎 0

package stringrange

//leetcode submit region begin(Prohibit modification and deletion)
func GenerateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	n1 := GenerateParenthesis(n - 1)
	retN := make([]string, 0)
	exists := make(map[string]struct{}, 0)
	for _, r := range n1 {
		key := "()" + r
		if _, ok := exists[key]; !ok {
			retN = append(retN, key)
			exists[key] = struct{}{}
		}
		for i := 0; i < len(r)-1; i++ {
			key = r[:i+1] + "()" + r[i+1:]
			if _, ok := exists[key]; !ok {
				retN = append(retN, key)
				exists[key] = struct{}{}
			}
		}
	}
	return retN
}

//leetcode submit region end(Prohibit modification and deletion)
