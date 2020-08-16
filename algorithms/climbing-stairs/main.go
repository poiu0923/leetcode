package main

import "fmt"

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	current, previous := 2, 1
	for i := 3; i <= n; i++ {
		current, previous = current+previous, current
	}
	return current
}

func main() {
	var n int

	n = 1
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))

	n = 2
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))

	n = 3
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))

	n = 10
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))

	n = 15
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))

	n = 44
	fmt.Printf("n = %d, result = %d\n", n, climbStairs(n))
}
