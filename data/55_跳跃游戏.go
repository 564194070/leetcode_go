package data

//自己一堆判断的麻烦解法
func canJump(nums []int) bool {
	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}

	if 1 == len(nums) {
		return true
	} else if 2 == len(nums) {
		if nums[0] > 0 {
			return true
		}
		return false
	}
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = max(dp[0]-1, nums[1])
	if dp[0] == 0 && 0 != len(nums)-1 {
		return false
	}
	if dp[1] == 0 && 1 != len(nums)-1 {
		return false
	}

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1]-1, nums[i])
		if dp[i] == 0 && i != len(nums)-1 {
			return false
		}
	}
	return true

}

//使用贪心算法优化了过程
func canJump2(nums []int) bool {
	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	if 0 == nums[0] {
		if len(nums) == 1 {
			return true
		}
		return false
	}
	mostlen := 1
	for i := 0; i < len(nums); i++ {
		if mostlen > 0 {
			mostlen = max(mostlen-1, nums[i])
			if mostlen >= len(nums)-i-1 {
				return true
			}
		}
	}
	return false
}
