package data

//滚动数组解法
func climbStairs(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	var1, var2 := 1, 2
	for i := 3; i <= n; i++ {
		var1, var2 = var2, var1+var2
	}
	return var2

}

//矩阵快速幂
//通项公式
