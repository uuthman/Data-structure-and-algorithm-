package array

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	//SORTING THE SLICE IN ASCENDING ORDER
	sort.Ints(array)
	
	res := [][]int{}
	
	
	for i := 0; i < len(array); i++{
		left := i + 1;
		right := len(array) - 1
		for left < right{
			
			sum := array[i] + array[left] + array[right]
			if sum == target{
				var temp []int
				temp = append(temp,array[i],array[left],array[right])
				res = append(res,temp)
				left += 1
				right -=1
			}else if sum < target{
				left += 1
			}else if sum > target{
				right -=1
			}
		}
} 
	
	return res
}