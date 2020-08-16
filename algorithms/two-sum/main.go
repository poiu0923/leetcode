package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		other := target - v
		if otherI, ok := numMap[other]; ok {
			return []int{otherI, i}
		}

		numMap[v] = i
	}
	return nil
}

func main() {
	var nums []int
	var target int

	nums = []int{2, 7, 11, 15}
	target = 9
	fmt.Printf("nums=%v, target=%d, result=%v \n", nums, target, twoSum(nums, target))

	nums = []int{0, 7, 11, 15}
	target = 9
	fmt.Printf("nums=%v, target=%d, result=%v \n", nums, target, twoSum(nums, target))

	nums = []int{2, 7, 11, 15}
	target = 26
	fmt.Printf("nums=%v, target=%d, result=%v \n", nums, target, twoSum(nums, target))

	nums = []int{2, 7, 11, 15}
	target = 17
	fmt.Printf("nums=%v, target=%d, result=%v \n", nums, target, twoSum(nums, target))

	nums = []int{-3, 4, 3, 90}
	target = 0
	fmt.Printf("nums=%v, target=%d, result=%v \n", nums, target, twoSum(nums, target))
}
