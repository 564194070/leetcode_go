package data

import "strconv"

//自己想的解法 一轮便利处理bulls
//map记录非bull 二轮便利map处理cows

func getHint(secret string, guess string) string {

	bulls, cows := 0, 0
	map1 := make(map[byte]int, 10)
	map2 := make(map[byte]int, 10)
	for index := 0; index < len(secret); index++ {
		if secret[index] == guess[index] {
			bulls++
		} else {
			map1[guess[index]]++
			map2[secret[index]]++
		}
	}
	for key, _ := range map1 {
		for map1[key] > 0 && map2[key] > 0 {
			cows++
			map1[key]--
			map2[key]--
		}
	}
	var str string = strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return str
}

//解法升级 二轮便利不O(n)了 直接相加
func getHint2(secret string, guess string) string {

	bulls, cows := 0, 0
	array1, array2 := [10]int{0}, [10]int{0}

	for index := 0; index < len(secret); index++ {
		if secret[index] == guess[index] {
			bulls++
		} else {
			array1[guess[index]-'0']++
			array2[secret[index]-'0']++
		}
	}
	min := func(num1 int, num2 int) int {
		if num1 > num2 {
			return num2
		}
		return num1
	}

	for index, _ := range array1 {
		cows += min(array1[index], array2[index])
	}
	var str string = strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
	return str
}
