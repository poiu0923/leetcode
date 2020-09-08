package main

import "fmt"

/**
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
*/
func countSubstrings(s string) int {
	length := len(s)
	count := 0

	for i := 0; i < 2*length-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < length && s[l] == s[r] {
			l--
			r++
			count++
		}
	}

	return count
}

func main() {
	var s string

	s = "aaa"
	fmt.Printf("s = %s, count = %d \n", s, countSubstrings(s))

	s = "abc"
	fmt.Printf("s = %s, count = %d \n", s, countSubstrings(s))

	s = "aakdjiwue"
	fmt.Printf("s = %s, count = %d \n", s, countSubstrings(s))
}
