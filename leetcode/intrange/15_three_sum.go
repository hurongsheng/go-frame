//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针 
// 👍 2540 👎 0

package intrange

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)
	nums = quickSort(nums)
	for start := 0; start < len(nums) && nums[start] <= 0; start++ {
		if start != 0 && nums[start] == nums[start-1] {
			continue
		}
		for end := len(nums) - 1; end > start && nums[end] >= 0; end-- {
			if end != len(nums)-1 && nums[end] == nums[end+1] {
				continue
			}
			mid := 0 - nums[start] - nums[end]
			for i := end - 1; i > start; i-- {
				fmt.Printf("nums:%v,start:%v,i:%v,end:%v,mid:%v \n", nums, start, i, end, mid)
				if i != end-1 && nums[i] == nums[i+1] {
					continue
				}
				if nums[i] == mid {
					fmt.Printf("append:  nums:%v,start:%v,i:%v,end:%v\n", nums, nums[start], nums[i], nums[end])
					res = append(res, []int{nums[start], nums[i], nums[end]})
				}
				if nums[i] < mid {
					break
				}
			}

		}
	}
	return res
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
