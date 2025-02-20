package logic02

func Soal06B (slice [][]int) [][]int{
	start, n := 1, len(slice)
	for i:= range slice{
		for j:= range slice[i]{
			if (i % 2 == 0) {
				// // kalo di border +3, kalo di dalem border + 2
				slice[i][j] = start
				start += 3
				} else {
				slice[i][n-j-1] = start
				start += 2
			}
		}
	}

	return slice
}