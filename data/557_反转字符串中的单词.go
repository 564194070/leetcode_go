package data

func reverseWords(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		//判断每个单词的起始位置
		start := i
		for i < length && s[i] != ' ' {
			//将i移动到第一个空格的位置
			i++
		}
		for p := start; p < i; p++ {
			//将字符串反转加入ret中
			ret = append(ret, s[start+i-1-p])
		}
		for i < length && s[i] == ' ' {
			//将空格补上
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
