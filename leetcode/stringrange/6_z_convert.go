//将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。 
//
// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下： 
//
// L   C   I   R
//E T O E S I I G
//E   D   H   N
// 
//
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。 
//
// 请你实现这个将字符串进行指定行数变换的函数： 
//
// string convert(string s, int numRows); 
//
// 示例 1: 
//
// 输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
// 
//
// 示例 2: 
//
// 输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G 
// Related Topics 字符串 
// 👍 811 👎 0

package stringrange

//leetcode submit region begin(Prohibit modification and deletion)
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	zMap := make(map[int]string, 0)
	mod := 2*numRows - 2
	x, y := 0, 0
	for i := 0; i < len(s); i++ {
		//0(0,0) 1(0,1),2(0,2),3(0,3),4(1,2),5(2,1),6(3,0)
		zMap[y] = zMap[y] + s[i:i+1]
		if i%mod < numRows-1 {
			y++
		} else {
			x++
			y--
		}
	}
	ret := ""
	for i := 0; i < numRows; i++ {
		ret += zMap[i]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
