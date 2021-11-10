package data

//打家劫舍 自己的dp备忘录 但是代码可读性差的解法
func rob(nums []int) int {

	max := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	lenNum := len(nums) - 1
	if lenNum == 0 {
		return nums[0]
	}

	dp := []int{}
	dp = append(dp, nums[0])
	dp = append(dp, max(nums[0], nums[1]))

	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(dp[i-2]+nums[i], dp[i-1]))
	}
	return dp[lenNum]

}

//滚动数组 压缩状态
func rob2(nums []int) int {

	max := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	lenNum := len(nums) - 1
	if lenNum == 0 {
		return nums[0]
	}

	arr := [3]int{}
	arr[0], arr[1], arr[2] = nums[0], max(nums[0], nums[1]), max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		arr[2] = max(arr[0]+nums[i], arr[1])
		arr[1], arr[0] = arr[2], arr[1]
	}
	return arr[2]

}
