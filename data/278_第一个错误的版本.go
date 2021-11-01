package data

func firstBadVersion(n int) int {

	low := 1
	high := n

	for high != low {
		mid := low + (high-low)/2

		if isBadVersion(mid) == true {
			//出现了问题的版本 问题已经出现向前搜索
			high = mid
		} else {
			//没有问题的版本向前追溯
			low = mid + 1
		}
	}
	return low
}

//模拟isBadVersion函数的作用
func isBadVersion(n int) bool {
	if n == 5 {
		return true
	}
	return false
}
