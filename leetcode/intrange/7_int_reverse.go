//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。 
//
// 示例 1: 
//
// 输入: 123
//输出: 321
// 
//
// 示例 2: 
//
// 输入: -123
//输出: -321
// 
//
// 示例 3: 
//
// 输入: 120
//输出: 21
// 
//
// 注意: 
//
// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231, 231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。 
// Related Topics 数学 
// 👍 2145 👎 0

package intrange

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	flag := true
	if x < 0 {
		flag = false
		x = -x
	}
	xSlice := make([]int, 0)
	for {
		xSlice = append(xSlice, x%10)
		if x >= 10 {
			x = x / 10
		} else {
			break
		}
	}
	x = 0
	for _, c := range xSlice {
		x = x*10 + c
	}

	if !flag {
		x = -x
	}
	if x > 1<<31-1 || x < -1<<31 {
		return 0
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
