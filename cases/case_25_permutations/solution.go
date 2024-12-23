package case_25_permutations

import "sort"

func permutations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var result [][]int
	var permute func(int)
	sort.Ints(nums)
	permute = func(start int) {
		if start == len(nums)-1 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}
		seen := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if seen[nums[i]] {
				continue // Skip duplicate elements
			}
			seen[nums[i]] = true
			nums[start], nums[i] = nums[i], nums[start]
			permute(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	permute(0)
	return result
}
