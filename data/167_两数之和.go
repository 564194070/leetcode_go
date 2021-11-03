package data

//使用二分查找的方法 查找有序数组中的两数之和
func twoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		low := 1 + i
		high := len(numbers) - 1
		temptarget := target - numbers[i]
		for low <= high {
			mid := low + (high-low)/2
			if numbers[mid] == temptarget {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > temptarget {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

//使用双指针方法 找到有序数组中的两数之和
func twoSumSecond(numbers []int, target int) []int {

	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}

}
