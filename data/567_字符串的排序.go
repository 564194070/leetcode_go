package data

//使用滑动窗口处理字符串找子序列问题

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}
	var arrs1, arrs2 [26]int

	for index, _ := range s1 {
		arrs1[s1[index]-'a']++
		arrs2[s2[index]-'a']++
	}
	if arrs1 == arrs2 {
		return true
	}

	for left, right := 0, len(s1); right < len(s2); left, right = left+1, right+1 {
		arrs2[s2[left]-'a']--
		arrs2[s2[right]-'a']++
		if arrs1 == arrs2 {
			return true
		}

	}
	return false
}
