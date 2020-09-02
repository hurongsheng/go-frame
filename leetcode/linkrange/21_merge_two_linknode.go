//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
// 
// Related Topics 链表 
// 👍 1242 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

package linkrange

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		return &ListNode{l1.Val,mergeTwoLists(l1.Next, l2)}
	} else {
		return &ListNode{l2.Val,mergeTwoLists(l1, l2.Next)}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
