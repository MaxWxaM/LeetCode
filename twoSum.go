package main

import "fmt"

func main() {
	result := twoSum([]int{1, 2, 3}, 5)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		if _, ok := resultMap[target-nums[i]]; ok {
			return []int{resultMap[target-nums[i]], i}
		}
		resultMap[nums[i]] = i
	}
	return nil
}
