package main

import (
	"fmt"
	. "frame/leetcode/linkrange"
	"frame/leetcode/stringrange"
)

func main() {
	//res := CanVisitAllRooms([][]int{{1}, {2}, {3}, {}})
	//res := TwoSum([]int{2, 11, 7, 5}, 9)
	//res := AddTwoNumbers(
	//
	//	&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	//)
	//print(RemoveNthFromEnd(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, 1), nil)
	//print(RemoveNthFromEnd(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, 2), nil)
	//print(RemoveNthFromEnd(&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, 3), nil)
	//print(LengthOfLongestSubstring("bbbbb"), 1)
	//print(LengthOfLongestSubstring("abb"), 2)
	//print(LengthOfLongestSubstring("abc"), 3)
	//print(LengthOfLongestSubstring("abba"), 2)
	//res := LongestPalindrome("abcbbcbb")
	//res := Convert("PAYPALISHIRING", 3)
	//res := StrToInt("9223372036854775808")
	//res := IsPalindrome(10)
	//res := IsMatch("aab", "c*a*b*")
	//res := IntToRoman(58)
	//res := LongestCommonPrefix([]string{"flower", "flow", "floght"})
	//print(intrange.ThreeSumClosest([]int{-1, 0, 1, 2, -1, -4}, 2), 2)
	//print(intrange.ThreeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
	print(stringrange.GenerateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})

}

func print(i interface{}, eq interface{}) {
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
		fmt.Printf("i  :%v\n", i)
		fmt.Printf("eq :%v\n", eq)
		fmt.Printf("res:%v\n", fmt.Sprintf("%+v", i) == fmt.Sprintf("%+v", eq))
	}
	fmt.Println("=========")
	fmt.Println("")
}
