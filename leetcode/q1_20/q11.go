package leet

func Q11MaxArea(height []int) int {
	n := len(height)
	ptrb, ptre := 0, n-1
	var res int
	for ptrb != ptre {
		smaller := min(height[ptrb], height[ptre])
		temp := smaller * (ptre - ptrb)
		if temp > res {
			res = temp
		}
		if smaller == height[ptrb] {
			ptrb++
		} else {
			ptre--
		}
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
