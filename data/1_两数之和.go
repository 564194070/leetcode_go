package data

//使用哈希表解决两数之和问题

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, num := range nums {
		num2, ok := m[target-num]

		if ok == true {
			return []int{num2, idx}
		}
		m[num] = idx
	}
	return nil
}
