package data

func sortedSquares(nums []int) []int {

	lastNum := -1
	n := len(nums)
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNum = i
	}

	m1 := make([]int, 0, n)
	for i, j := lastNum, lastNum+1; i >= 0 || j < n; {
		if i < 0 {
			m1 = append(m1, nums[j]*nums[j])
			j++
		} else if j == n {
			m1 = append(m1, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] > nums[j]*nums[j] {
			m1 = append(m1, nums[j]*nums[j])
			j++
		} else {
			m1 = append(m1, nums[i]*nums[i])
			i--
		}
	}
	return m1
}
