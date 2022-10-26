package easy

func (*Practice) FindLinkedListCycle(head *ListNode) bool {
	var p2 *ListNode
	var p = head

	if head != nil && head.Next != nil {
		p2 = head.Next
	}

	for p != nil && p2 != nil {
		if p == p2 {
			return true
		}
		p = p.Next
		if p2.Next != nil && p.Next.Next != nil {
			p2 = p2.Next.Next
		} else {
			p2 = nil
		}
	}
	return false
}
