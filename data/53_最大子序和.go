package data

//使用动态规划处理最大子序和问题
func maxSubArray(nums []int) int {

	max := func(num1 int, num2 int) int {
		if num1 < num2 {
			return num2
		}
		return num1
	}
	sum := nums[0]
	dp := []int{}
	dp = append(dp, nums[0])
	for i := 1; i < len(nums); i++ {
		dp = append(dp, max(dp[i-1]+nums[i], nums[i]))
		sum = max(sum, dp[i])
	}
	return sum
}
