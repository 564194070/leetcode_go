package data

//哈希表 是否存在重复元素
func containsDuplicate(nums []int) bool {
	map1 := make(map[int]int)

	for _, value := range nums {
		valuemap, exists := map1[value]
		if exists {
			if 1 == valuemap {
				return true
			}
		} else {
			map1[value] = 1
		}
	}
	return false
}
