package logic02

func Soal06A (slice [][]int) [][]int {
	start:= 1
	n := len(slice)
	for i:= range slice{
		for j:= range slice[i]{
			if(i%2==0){
				slice[i][j] = start
				start += 3
			} else {
				reverseCol := n - j - 1
				slice[i][reverseCol] = start
				start += 2
			}
		}
	}
	return slice
}