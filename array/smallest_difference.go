package array

import (
	"sort"
	"math"
) 

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.

	sort.Ints(array1)
	sort.Ints(array2)
	
	
	
	left1 := 0
	left2 := 0
	s := math.Inf(1)
	cal := math.Inf(1)
	sp := []int{}
	
	
	for left1 < len(array1) && left2 < len(array2){
		
		f := array1[left1]
		se := array2[left2]
		
		if f < se{
			cal = float64(se) - float64(f)
			left1 += 1
		}else if f > se {
			cal = float64(f) - float64(se)
			left2 += 1
		}else{
			return []int{f,se}
		}
		
		if s > cal{
			s = cal
			sp = []int{f,se}
		}
	}
	return sp
}