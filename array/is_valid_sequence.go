package array

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here
	a := 0
	s := 0
	for a < len(array) && s < len(sequence) {
		if array[a] == sequence[s]{
			s++
		}
		a++
	}
	return s == len(sequence)
	
}
