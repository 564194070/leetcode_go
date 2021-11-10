package data

//将这类问题 统一转化为打家劫舍问题 然后再处理

func deleteAndEarn(nums []int) int {

	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
	maxVal := 0

	for _, value := range nums {
		maxVal = max(maxVal, value)
	}

	arr := make([]int, maxVal+1)
	for _, value := range nums {
		arr[value] += value
	}

	left, right, temp := arr[0], max(arr[0], arr[1]), max(arr[0], arr[1])
	for i := 2; i < maxVal+1; i++ {
		temp = max(left+arr[i], right)
		left, right = right, temp
	}
	return temp
}
