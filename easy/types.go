package easy

type Practice struct {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	p := head
	for i := range arr[1:] {
		p.Next = &ListNode{Val: arr[i+1]}
		p = p.Next
	}
	return head
}
