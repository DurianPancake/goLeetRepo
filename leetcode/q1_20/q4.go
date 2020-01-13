package leet

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	left, right := (n+m+1)/2, (n+m+2)/2
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left)+
		getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return int(math.Min(float64(nums1[start1]), float64(nums2[start2])))
	}

	var i = start1 + int(math.Min(float64(len1), float64(k/2))) - 1
	var j = start2 + int(math.Min(float64(len2), float64(k/2))) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}
