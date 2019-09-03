package main

func twoSum(nums []int, target int) []int {
	args := make(map[int]int)
	res := make([]int, 2)
	for index, value := range nums {
		if v, ok := args[target-value]; ok {
			res[0] = index
			res[1] = v
			return res
		}
		args[value] = index
	}
	return res
}