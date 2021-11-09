package data

//不知道为啥不对的方法
func merge(nums1 []int, m int, nums2 []int, n int) {

	for num1, num2, index := m-1, n-1, m+n-1; num2 >= 0; {
		if num1 == -1 {
			nums1[index] = nums2[num2]
			num2--
		} else {
			if nums1[num1] < nums2[num2] {
				nums1[index] = nums2[num2]
				num2--
			} else {
				nums1[index] = nums1[num1]
				num1--
			}
		}
		index--

	}

}

//观察解体后的答案
func merge2(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
