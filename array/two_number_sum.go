package array

import "sort"

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	
	for left < right {
		
		sum := array[left] + array[right]
		
		if sum == target{
			return []int{array[left],array[right]}
		}else if sum < target{
			left++
		}else if sum > target{
			right--
		}	
	}
	return []int{}
}