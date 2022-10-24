package easy

type Stack struct {
	data []byte
}

func newStack() *Stack {
	s := &Stack{}
	return s
}
func (s *Stack) pop() byte {
	if s.len() == 0 {
		return 0
	}
	lastIdx := s.len() - 1
	last := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return last
}

func (s *Stack) push(v byte) {
	s.data = append(s.data, v)
}

func (s *Stack) len() int {
	return len(s.data)
}

func (p *Practice) ValidParentheses(s string) bool {
	stack := newStack()

	for i := range s {
		p := s[i]
		isOpen := false
		required := byte(0)

		switch p {
		case '{':
			isOpen = true
		case '}':
			required = '{'
		case '(':
			isOpen = true
		case ')':
			required = '('
		case '[':
			isOpen = true
		case ']':
			required = '['
		default:
			continue
		}
		if isOpen {
			stack.push(p)
			continue
		}

		if stack.len() == 0 || stack.pop() != required {
			return false
		}
	}

	return stack.len() == 0
}
