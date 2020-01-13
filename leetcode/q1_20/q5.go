package leet

import "math"

func LongestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}
	begin, maxLength := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLength/2 {
			break
		}
		b, e := i, i
		for e < len(s)-1 && s[e] == s[e+1] {
			e++
		}
		i = e + 1
		for e < len(s)-1 && b > 0 && s[b-1] == s[e+1] {
			e++
			b--
		}
		newLen := e + 1 - b
		// 创新记录的话，就更新记录
		if newLen > maxLength {
			begin = b
			maxLength = newLen
		}
	}
	return s[begin : begin+maxLength]
}

func L2(s string) string {
	runes := []rune(s)
	var maxStart, cursor uint
	var maxCenter, currCenter float32

	arrayLen := len(runes)
	for i := 0; i < 2*arrayLen-1; i++ {
		currCenter = float32(i) / 2
		cursor = uint(math.Floor(float64(i / 2)))
		for {
			crossPtr := uint(i) - cursor
			if isInRange(crossPtr, arrayLen) {
				if runes[cursor] == runes[crossPtr] {
					if maxCenter-float32(maxStart) < currCenter-float32(cursor) {
						maxStart, maxCenter = cursor, currCenter
					}
					if cursor > 0 {
						cursor--
						continue
					} else {
						break
					}
				} else {
					break
				}
			} else {
				// return
				cross := int(2*maxCenter) - int(maxStart)
				return string(runes[maxStart : cross+1])
			}
		}
	}
	return s
}

func isInRange(ptr uint, len int) bool {
	return ptr < uint(len) && ptr >= 0
}
