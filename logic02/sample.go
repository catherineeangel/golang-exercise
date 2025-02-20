package logic02

import slice "github.com/catherineeangel/go-print-slice"

func Sample02 (){
	// init slice
	// n := 10
	slice2D := initMatrix(5)
	
	Soal07(&slice2D)
	slice.Print2DSlice(slice2D)

	slice2D = initMatrix(5)
	Soal08(&slice2D)
	slice.Print2DSlice(slice2D)

	Soal09(&slice2D)
	slice.Print2DSlice(slice2D)
}

func initMatrix(n int) [][]int{
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return matrix
}