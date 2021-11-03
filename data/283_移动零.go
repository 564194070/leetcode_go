package data

//双指针处理问题
func moveZeroes(nums []int) {

	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
