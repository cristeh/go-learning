package algos

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	sort.Ints(nums) //Sorting greatly simplifies the code since it's gonna be n^2 anyway
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { //Can't have a, b, c = 0 if a, b, c > 0
			break
		}
		if i == 0 || nums[i-1] != nums[i] { //Deduplicate; do Two Sum problem
			j, k := i+1, len(nums)-1
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 { //We have found our pair!
					result = append(result, []int{nums[i], nums[j], nums[k]})
					j++
					k--
					for j < k && nums[j-1] == nums[j] { //Deduplicate
						j++
					}
				} else if sum < 0 { //Sum needs to be bigger, increase j
					j++
				} else { //Sum needs to be smaller, decrease k
					k--
				}
			}
		}
	}
	return result
}
