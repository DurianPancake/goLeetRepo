package leet

func Convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	var term = numRows*2 - 2
	length := len(s)
	runes := make([]uint8, length, length)
	// i is row and cursor is for natural growth order count
	for i, cursor := 0, 0; i < numRows; i++ {
		// offset is big period movement
		for offset := 0; offset+i < length; offset += term {
			runes[cursor] = s[i+offset]
			cursor++
			// when the row in the middle rows has two num in the big period
			if i != 0 && i != numRows-1 && offset+term-i < length {
				runes[cursor] = s[offset+term-i]
				cursor++
			}
		}
	}
	return string(runes)
}
func Convert2(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	rowIndex := 0
	goDown := false
	for i := 0; i < len(s); i++ {
		rows[rowIndex] = append(rows[rowIndex], s[i])
		if rowIndex == 0 || rowIndex == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			rowIndex++
		} else {
			rowIndex--
		}
	}
	var resStr string
	for i := 0; i < len(rows); i++ {
		resStr += string(rows[i])
	}
	return resStr
}
