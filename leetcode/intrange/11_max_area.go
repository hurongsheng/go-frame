//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, 
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。 
//
// 
//
// 
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。 
//
// 
//
// 示例： 
//
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49 
// Related Topics 数组 双指针 
// 👍 1784 👎 0

package intrange

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxArea := 0
	for {
		area := min(height[start], height[end]) * (end - start)
		if maxArea < area {
			maxArea = area
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
		if start >= end {
			break
		}
	}
	return maxArea
}

func maxArea2(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := min(height[j], height[i]) * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
