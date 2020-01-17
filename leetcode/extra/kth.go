package extra

import "math"

func NthUglyNumber(n int, a int, b int, c int) int {
	return getKth(a, b, c, 0, 0, 0, n)
}

func getKth(a int, b int, c int, m int, n int, o int, k int) int {
	if k == 1 {
		return getMin(a*(m+1), b*(n+1), c*(o+1))
	}
	v1, v2, v3 := a*m+a*(k/2), b*n+b*(k/2), c*o+c*(k/2)
	min := getMin(v1, v2, v3)
	if v1 == min {
		offset := getOffset(min, b, c, n, o)
		return getKth(a, b, c, m+k/2, n, o, k-k/2+offset)
	} else if v2 == min {
		offset := getOffset(min, a, c, m, o)
		return getKth(a, b, c, m, n+k/2, o, k-k/2+offset)
	} else {
		offset := getOffset(min, b, a, n, m)
		return getKth(a, b, c, m, n, o+k/2, k-k/2+offset)
	}
}

func getOffset(big int, f1 int, f2 int, n1 int, n2 int) int {
	min := int(math.Min(float64(f1*n1), float64(f2*n2)))
	lcm := getLcm(f1, f2)
	return big/lcm - (min-1)/lcm
}

func getLcm(f1 int, f2 int) int {
	return (f1 * f2) / getGcd(f1, f2)
}

func getGcd(f1 int, f2 int) int {
	temp := f1 % f2
	if temp > 0 {
		return getGcd(f2, temp)
	}
	return f2
}

func getMin(a int, b int, c int) int {
	temp := math.Min(float64(a), float64(b))
	return int(math.Min(temp, float64(c)))
}
