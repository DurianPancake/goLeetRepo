package leet

func Q7Reverse(x int) int {

	const (
		intMax = int(^uint32(0) >> 1)
		intMin = ^intMax
	)
	var i int
	for x != 0 {
		i = i*10 + x%10
		x /= 10
	}

	if i < intMax && i > intMin {
		return i
	}
	return 0
}
