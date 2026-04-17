/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Next == nil {
        return nil
    }

    slow, fast := head, head
    var temp *ListNode
    temp = nil

    for fast != nil && fast.Next != nil {
        temp = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    if slow != nil {
        temp.Next = slow.Next
    }
    return head
}