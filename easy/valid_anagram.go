package arraynhashing

func (p *Practice) IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		chars[s[i]-byte('a')]--
		chars[t[i]-byte('a')]++
	}

	for i := 0; i < 26; i++ {
		if chars[byte(i)] != 0 {
			return false
		}
	}

	return true
}
