package data

//可以运行 但是时间复杂度高 测试不通过
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		temp := nums[0]
		nums[0] = nums[len(nums)-1]
		for j := 0; j < len(nums)-1; j++ {
			temp, nums[j+1] = nums[j+1], temp
		}
	}
}

//反转数组
func rotate2(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
