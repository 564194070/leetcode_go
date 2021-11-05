package data

//面对动态规划问题的三个要素
//重叠子问题，最优子结构，状态转移方程

//暴力解
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//自顶向下
func fib2(n int) int {
	if n < 1 {
		return 0
	}

	return helper(n)
}

func helper(n int) int {
	map1 := make(map[int]int, n)
	if n == 1 || n == 2 {
		return 1
	}
	if value, ok := map1[n]; ok != false {
		return value
	} else {
		return helper(n-1) + helper(n-2)
	}
}

//自底向上
func fib3(n int) int {
	map1 := make(map[int]int, n)
	map1[0], map1[1], map1[2] = 0, 1, 1

	for i := 3; i <= n; i++ {
		map1[i] = map1[i-1] + map1[i-2]
	}
	return map1[n]
}

//状态压缩
//发现斐波那契数列只需要两个状态 所以维护两个状态即可
func fib4(n int) int {
	var nums int
	if n < 1 {
		return 0
	} else if n < 3 {
		return 1
	} else {
		left, right := 0, 1
		for i := 2; i <= n; i++ {
			nums = left + right
			left, right = right, left+right
		}
		return nums
	}
}
