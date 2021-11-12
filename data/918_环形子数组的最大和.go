package data

//环形子数组的最大和 迄今为止发现的很难理解的题目
func maxSubarraySumCircular(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	Maxtemp, Mintemp, Sum, Min, Max := nums[0], nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		Maxtemp = max(Maxtemp+nums[i], nums[i])
		Mintemp = min(Mintemp+nums[i], nums[i])
		Max = max(Max, Maxtemp)
		Min = min(Min, Mintemp)
		Sum += nums[i]
	}

	if Max < 0 {
		return Max
	}
	return max(Sum-Min, Max)
}
