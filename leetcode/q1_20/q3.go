package leet

import "strings"

func LengthOfLongestSubstring(s string) int {

	var start, end, max int
	runes := []rune(s)

	// main loop
main:
	for i, r := range runes {
		window := runes[start:end]
		for j, rw := range window {
			// if i and j point to the same position
			if j+start == i {
				continue
			}
			if r == rw {
				// compute window's size
				if tempSize := end - start; tempSize > max {
					max = tempSize
				}
				// move the window
				start += j + 1 // j 's next position
				end++
				continue main
			}
		}
		// expand window
		end++
	}
	if tempSize := end - start; tempSize > max {
		max = tempSize
	}
	return max
}

// 不能兼容中文
func LenM(s string) int {
	start, window := 0, 0
	for key := 0; key < len(s); key++ {
		posIndex := strings.Index(s[start:key], string(s[key]))
		if posIndex == -1 {
			if key-start+1 > window {
				window = key - start + 1
			}
		} else {
			start = start + 1 + posIndex
		}
	}
	return window
}
