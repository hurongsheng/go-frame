//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。 
//
// 示例： 
//
// 给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
// 
//
// 说明： 
//
// 给定的 n 保证是有效的。 
//
// 进阶： 
//
// 你能尝试使用一趟扫描实现吗？ 
// Related Topics 链表 双指针 
// 👍 970 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package linkrange

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	ret, point := head, head
	m := 0
	for point.Next != nil {
		point = point.Next
		if m >= n {
			head = head.Next
		}
		m++
	}
	if m < n {
		return ret.Next
	}
	head.Next = head.Next.Next
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
