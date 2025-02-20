package logic01

// 1
func OddPattern(n int) []int {
	slice := make([]int, n)
	for i:= range n {
		slice[i] = 2 * i + 1
	}
	return slice
}

// 2
func EvenPattern(n int) []int {
	slice := make([]int, n)
	for i:= range n {
		slice[i] = 2 * (i + 1)
	}
	return slice
}

// 3
func MultipleThreePattern(n int) []int {
	slice := make([]int, n)
	for i:= range n {
		slice[i] = 3 * (i + 1)
	}
	return slice
}