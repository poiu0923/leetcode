package main

import "fmt"

/**
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.

注意：
输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-consecutive-ones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMaxConsecutiveOnes(nums []int) int {
	maxConsecutive := 0
	consecutive := 0
	for i, v := range nums {
		if v == 1 {
			consecutive++
		}
		if i == len(nums)-1 || v != 1 {
			if consecutive > maxConsecutive {
				maxConsecutive = consecutive
			}
			consecutive = 0
		}
	}

	return maxConsecutive
}

func main() {
	var nums []int

	nums = []int{1, 1, 0, 1, 1, 1}
	fmt.Printf("nums = %v, maxConsecutive = %d \n", nums, findMaxConsecutiveOnes(nums))

	nums = []int{1, 1, 1, 1, 1, 1}
	fmt.Printf("nums = %v, maxConsecutive = %d \n", nums, findMaxConsecutiveOnes(nums))

	nums = []int{0, 1, 0, 0, 1, 1}
	fmt.Printf("nums = %v, maxConsecutive = %d \n", nums, findMaxConsecutiveOnes(nums))

	nums = []int{0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0}
	fmt.Printf("nums = %v, maxConsecutive = %d \n", nums, findMaxConsecutiveOnes(nums))
}