package main

import "fmt"

/**
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:

输入: [3, 2, 1]

输出: 1

解释: 第三大的数是 1.
示例 2:

输入: [1, 2]

输出: 2

解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:

输入: [2, 2, 3, 1]

输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/third-maximum-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func thirdMax(nums []int) int {
	topMap := make(map[int]int)
	for _, v := range nums {

		if m1, ok := topMap[1]; ok {
			if v > m1 {
				topMap[1], topMap[2], topMap[3] = v, topMap[1], topMap[2]
				continue
			}
		} else {
			topMap[1] = v
			continue
		}

		if m2, ok := topMap[2]; ok {
			if v > m2 {
				topMap[2], topMap[3] = v, topMap[2]
				continue
			}
		} else {
			topMap[2] = v
			continue
		}

		if m3, ok := topMap[3]; ok {
			if v > m3 {
				topMap[3] = v
			}
		} else {
			topMap[3] = v
		}
	}

	if len(topMap) > 2 {
		return topMap[3]
	} else {
		return topMap[1]
	}
}

func main() {
	fmt.Println(thirdMax([]int{0, 0, 0, 0}))
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}
