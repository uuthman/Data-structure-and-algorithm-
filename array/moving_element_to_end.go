package array

func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	left := 0
	right := len(array) - 1
	
	
	for left < right{
	
		for left < right &&  array[right] == toMove{
			right--
		}
		
		if array[left] == toMove {
			array[left],array[right] = array[right],array[left]
		}
		left++
	}
	return array
}