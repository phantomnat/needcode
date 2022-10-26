package easy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

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

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprintf("%d", l.Val))
	mem := make(map[string]struct{})
	mem[fmt.Sprintf("%p", l)] = struct{}{}

	p := l.Next
	for p != nil {
		key := fmt.Sprintf("%p", p)
		if _, ok := mem[key]; ok {
			sb.WriteString(" -> (cycle detected)")
			return sb.String()
		}
		mem[key] = struct{}{}
		sb.WriteString(fmt.Sprintf(" -> %d", p.Val))
		p = p.Next
	}
	return sb.String()
}

var _ json.Marshaler = (*ListNode)(nil)

func (l *ListNode) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(struct {
		Out string `json:"out"`
	}{
		Out: l.String(),
	})
	return buf.Bytes(), nil
	// return json.Marshal(struct {
	// 	Out string `json:"out"`
	// }{
	// 	Out: l.String(),
	// })
}
