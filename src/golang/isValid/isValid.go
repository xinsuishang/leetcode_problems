package isValid

func isValid(s string) bool {
	var c []byte
	symbol := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, value := range s {
		clen := len(c)
		if clen > 0 {
			if _, ok := symbol[byte(value)]; ok {
				if c[clen-1] == symbol[byte(value)] {
					c = c[:clen-1]
					continue
				} else {
					return false
				}
			}
		}
		c = append(c, byte(value))
	}
	return len(c) == 0
}

// times: 1
