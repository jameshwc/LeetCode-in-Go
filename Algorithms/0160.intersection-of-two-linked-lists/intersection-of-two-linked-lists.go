package problem0160

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	hasLinkedToB, hasLinkedToA := false, false
	for a != nil && b != nil {
		if a == b {
			return b
		}
		a, b = a.Next, b.Next
		if a == nil && !hasLinkedToB {
			hasLinkedToB = true
			a = headB
		}
		if b == nil && !hasLinkedToA {
			hasLinkedToA = true
			b = headA
		}
	}
	return nil
}
