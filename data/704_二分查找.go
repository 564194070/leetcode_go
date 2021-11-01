package data

//不使用二分查找 直接采用穷举法
func search(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}
	return -1
}

//使用二分查找
func binarysearch(nums []int, target int) int {

	high := len(nums) - 1
	low := 0

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid

		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
