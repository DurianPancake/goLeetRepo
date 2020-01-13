package leet

func Q9IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	x0, reverse := x, 0
	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == x0
}
