//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。 
//
// 
//
// 示例： 
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 
//
// 
//
// 提示： 
//
// 
// 3 <= nums.length <= 10^3 
// -10^3 <= nums[i] <= 10^3 
// -10^4 <= target <= 10^4 
// 
// Related Topics 数组 双指针 
// 👍 560 👎 0
package intrange

import "sort"

//import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var closest *int
	closestNum := 0
	for start := 0; start < len(nums); start++ {
		if start != 0 && nums[start] == nums[start-1] {
			continue
		}
		for end := len(nums) - 1; end > start; end-- {
			if end != len(nums)-1 && nums[end] == nums[end+1] {
				continue
			}
			for i := end - 1; i > start; i-- {
				if i != end-1 && nums[i] == nums[i+1] {
					continue
				}
				s := nums[start] + nums[i] + nums[end]
				closer := abs(s, target)
				if closer == 0 {
					return target
				}
				if closest == nil {
					closest = &s
					closestNum = closer
				}
				if closer < closestNum {
					closest = &s
					closestNum = closer
				}
			}
		}
	}
	return *closest
}
func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

//leetcode submit region end(Prohibit modification and deletion)
