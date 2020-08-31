//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。 
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。 
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。 
//
// 示例： 
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
// 
// Related Topics 链表 数学 
// 👍 4821 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package linkrange

//
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return AddTwoNumbers2(l1, l2, 0)
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode, flag int) *ListNode {
	if l1 == nil && l2 == nil {
		if flag == 1 {
			return &ListNode{1, nil}
		}
		return nil

	}
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	add := l1.Val + l2.Val + flag
	return &ListNode{
		Val:  add % 10,
		Next: AddTwoNumbers2(l1.Next, l2.Next, add/10),
	}
}

//leetcode submit region end(Prohibit modification and deletion)
