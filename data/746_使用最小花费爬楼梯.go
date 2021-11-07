package data

//使用dp数组备忘录解法
func minCostClimbingStairs(cost []int) int {

	high := len(cost)
	dp := make(map[int]int, high+1)

	dp[0] = 0
	dp[1] = 0
	min := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num2
		}
		return num1
	}

	for i := 2; i <= high; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[high]
}

//压缩状态 滚动数组
func minCostClimbingStairs2(cost []int) int {

	high := len(cost)
	dp1, dp2, temp := 0, 0, 0
	min := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num2
		}
		return num1
	}

	for i := 2; i <= high; i++ {
		temp = min(dp1+cost[i-1], dp2+cost[i-2])
		dp2, dp1 = dp1, temp
	}
	return temp
}
