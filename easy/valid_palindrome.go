package easy

func (p *Practice) IsPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	isChar := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z')
	}
	toLower := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c + 'a' - 'A'
		}
		return c
	}
	isEqual := func(a, b byte) bool {
		return toLower(a) == toLower(b)
	}

	// O(N)
	for l < r && l < len(s) && r >= 0 {
		if !isChar(s[l]) {
			l++
			continue
		} else if !isChar(s[r]) {
			r--
			continue
		}

		if !isEqual(s[l], s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
