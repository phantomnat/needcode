package easy

func (*Practice) ReverseLinkedList(head *ListNode) *ListNode {
	var mem []*ListNode
	if head == nil {
		return head
	}
	ptr := head
	for ptr != nil {
		mem = append(mem, ptr)
		ptr = ptr.Next
	}
	for i := len(mem)-1; i > 0; i-- {
		mem[i].Next = mem[i-1]
	}
	mem[0].Next = nil
	return mem[len(mem)-1]
}
