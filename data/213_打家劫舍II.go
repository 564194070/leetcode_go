package data

//使用两个dp处理两个 不同的 带头尾的计数
func rob(nums []int) int {
	max := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}

	houselen := len(nums)
	if houselen == 1 {
		return nums[0]
	} else if houselen == 2 {
		return max(nums[0], nums[1])
	}

	dp, dp1 := []int{}, []int{}
	dp, dp1 = append(dp, nums[0]), append(dp1, nums[1])
	dp, dp1 = append(dp, max(nums[0], nums[1])), append(dp1, max(nums[1], nums[2]))

	for i := 2; i < houselen-1; i++ {
		dp = append(dp, max(dp[i-2]+nums[i], dp[i-1]))
	}

	for i := 3; i < houselen; i++ {
		dp1 = append(dp1, max(dp1[i-3]+nums[i], dp1[i-2]))
	}

	return max(dp[houselen-2], dp1[houselen-2])
}
