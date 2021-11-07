package data

//使用状态压缩方法 处理动态规划的泰伯那契数
//压缩状态和滚动数组
func tribonacci(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	temp1, temp2, temp3 := 1, 1, 2

	for i := 4; i <= n; i++ {
		temp1, temp2, temp3 = temp2, temp3, temp1+temp2+temp3
	}
	return temp3
}
