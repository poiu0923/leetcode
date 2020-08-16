package main

import (
	"fmt"
	"strconv"
)

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	length := len(str)
	for i, ch := range str {
		if otherI := length - 1 - i; otherI >= i && uint8(ch) != str[otherI] {
			return false
		}
	}

	return true
}

func main() {
	x := 0
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))

	x = 123
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))

	x = 12343422
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))
	x = 123434220
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))

	x = -121
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))

	x = 121
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))

	x = 123321
	fmt.Printf("x=%d, result=%v \n", x, isPalindrome(x))
}
