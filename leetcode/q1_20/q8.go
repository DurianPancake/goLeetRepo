package leet

const (
	intMax = int(^uint32(0) >> 1)
	intMin = ^intMax
)

func Q8MyAtoi(s string) int {

	// trim ' '
	nega := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if nega = s[i] == '-'; nega || s[i] == '+' {
				i++
			}
			s = s[i:]
			break
		}
	}

	var result int
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			break
		}
		result = result*10 + int(ch)
		if result > intMax {
			result = intMax
			if nega {
				result = ^result
			}
			return result
		}
	}
	if nega {
		result = -result
	}
	return result
}
