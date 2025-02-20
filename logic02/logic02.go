package logic02

import (
	"fmt"

	slice "github.com/catherineeangel/go-print-slice"
	"github.com/catherineeangel/golang-exercise/logic01"
)


func Soal01(slice *[][]int){
	fmt.Println("Odd Patterns")
	for i := range *slice{
 		logic01.OddPattern(&((*slice)[i]))
		// for j := range (*slice)[i]{
		// 	(*slice)[i][j] = (2 * j) + 1	
		// }
	}

}

func Soal01Alt(n int) (result [][]int) {
	slice := slice.CreateSlice(n)

	// fmt.Println("Odd Patterns")
	for i := range slice{
 		logic01.OddPattern(&((slice)[i]))
		// for j := range (*slice)[i]{
		// 	(*slice)[i][j] = (2 * j) + 1	
		// }
	}
	return slice

}
func Soal02(slice *[][]int){
	fmt.Println("Even Patterns")
	for i := range *slice{
		logic01.EvenPattern(&((*slice)[i]))
		// for j := range (*slice)[i]{
		// 	(*slice)[i][j] = 2 * (j + 1)	
		// }
	}

}

func Soal03(slice *[][]int){
	fmt.Println("Odd Continuous Patterns")

	start:= 1
	for i := range *slice{
		for j:= range (*slice)[i]{
			(*slice)[i][j] = start
			start += 2
		}
	}

}

func Soal06(slice *[][]int){
	/* 
	row genap +3
	row ganjil +2
	*/
	n := len(*slice)
	fmt.Println("Plus 3 & 2 Patterns")
	start := 1
	for i := range *slice{
		for j := range (*slice)[i]{
			switch {
			case i % 2 == 0 :
				(*slice)[i][j] = start
				start += 3
			default:
				col := n - 1 - j
				(*slice)[i][col] = start
				start += 2
				
			}
		}
	}

}


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

func Soal10(slice *[][]int){
	// segitiga kiri
	fmt.Println("Segitiga Kiri")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (j <= i){
				(*slice)[i][j] = (2 * j) + 1
			}

		}
	}

}

func Soal11(slice *[][]int){
	// segitiga kanan
	fmt.Println("Segitiga Kanan")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (i <= j){
				(*slice)[i][j] = (2 * j) + 1
			}

		}
	}

}

func Soal12(slice *[][]int){
	// segitiga kiri + kanan
	
	n := len(*slice)
	fmt.Println("Segitiga Kiri Kanan")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (j <= i && (j + i) < n){
				(*slice)[i][j] = (2 * j) + 1
			} else if (i <= j && (j + i) >= n - 1) {
				(*slice)[i][j] = (2 * j) + 1
			}

		}
	}

}

func Soal13(slice *[][]int){
	// diagonal 2
	n := len(*slice)
	fmt.Println("Hourglass Pattern")
	for i := range *slice{
		for j := range (*slice)[i]{
			if (i <= j && (j + i) < n){
				(*slice)[i][j] = (2 * j) + 1
			} else if (j <= i && (j + i) >= n - 1) {
				(*slice)[i][j] = (2 * j) + 1
			}

		}
	}
	// for i := range *slice{
	// 	for j := range (*slice)[i]{
	// 		if (i == j){
	// 			(*slice)[i][j] = (2 * i) + 1
	// 		} else if (j + i == len(*slice) - 1) {
	// 			(*slice)[j][i] = (2 * i) + 1
	// 		} else if (i < n/2) {
	// 			switch {
	// 			case  j >= i:
	// 				(*slice)[i][j] = (2 * j) + 1
	// 			}
	// 		}

	// 	}
	// }

}