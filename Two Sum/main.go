package main

func twoSum(nums []int, target int) []int {
	var res []int

	var size int = len(nums) - 1

	for i := 0; i < size; i++ {
		for j := i + 1; j <= size; j++ {
			if nums[i] + nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}

	return res
}