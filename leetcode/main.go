package main

import (
	"fmt"
	. "frame/leetcode/intrange"
	. "frame/leetcode/linkrange"
)

func main() {
	//res := CanVisitAllRooms([][]int{{1}, {2}, {3}, {}})
	//res := TwoSum([]int{2, 11, 7, 5}, 9)
	//res := AddTwoNumbers(
	//	&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
	//	&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	//)
	//res := LengthOfLongestSubstring("abcabcbb")
	//res := LongestPalindrome("abcbbcbb")
	//res := Convert("PAYPALISHIRING", 3)
	//res := StrToInt("9223372036854775808")
	//res := IsPalindrome(10)
	//res := IsMatch("aab", "c*a*b*")
	res := IntToRoman(58)
	print(res)

}

func print(i interface{}) {
	if l, ok := i.(*ListNode); ok {
		for {
			fmt.Println(l)
			if l.Next != nil {
				l = l.Next
			} else {
				break
			}

		}
	} else {
		fmt.Println("res:", i)
	}
	fmt.Println("=========")
	fmt.Println("")
}
