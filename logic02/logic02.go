package logic02

import (
	"fmt"
)

func Soal07(slice *[][]int){
	// diagonal 1
	fmt.Println("Diagonal")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (i == j){
				(*slice)[i][j] = (2 * i) + 1
			}

		}
	}

}
func Soal08(slice *[][]int){
	// diagonal 1
	fmt.Println("Diagonal")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (j + i == len(*slice) - 1){
				(*slice)[j][i] = (2 * i) + 1
			}

		}
	}

}
func Soal09(slice *[][]int){
	// diagonal 2
	fmt.Println("Diagonal 2")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (i == j){
				(*slice)[i][j] = (2 * i) + 1
			} else if (j + i == len(*slice) - 1) {
				(*slice)[j][i] = (2 * i) + 1
			}

		}
	}

}

func Soal13(slice *[][]int){
	// diagonal 2
	fmt.Println("Hourglass Pattern")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (i == j){
				(*slice)[i][j] = (2 * i) + 1
			} else if (j + i == len(*slice) - 1) {
				(*slice)[j][i] = (2 * i) + 1
			}

		}
	}

}