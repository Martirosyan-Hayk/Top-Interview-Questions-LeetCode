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

// solution 2
// func twoSum(nums []int, target int) []int {
//     numbersToFind := make(map[int]int);
    
//     for i := 0; i < len(nums); i++ {
//        numberToFind := target - nums[i];
//        value, found := numbersToFind[numberToFind];
//         if found  {
//             return []int{value, i};
//         }
//         numbersToFind[nums[i]] = i;
//     } 

//     return []int{};

// }