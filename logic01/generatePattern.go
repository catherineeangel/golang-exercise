package logic01

import (
	"fmt"
)

// Soal 1
func OddPattern(slice *[]int) {
	for i:= range *slice {
		(*slice)[i] = 2 * i + 1
	}
}

// Soal 2
func EvenPattern(slice *[]int) {
	for i:= range *slice {
		(*slice)[i] = 2 * (i + 1)
	}
}

// Soal 3
func MultipleThreePattern(slice *[]int) {
	for i:= range *slice {
		(*slice)[i] = 3 * (i + 1)
	}
}

// Soal 4
func ReverseOddPattern (slice *[]int) {
	n := len(*slice)
	for i:= range (*slice) {
		(*slice)[i] = (2 * (n - i - 1) + 1)
	}
}

// Soal 7
func RepeatOddPattern (slice *[]int){
	n := len(*slice)
	for i:= range (*slice){
		if (i < n/2){
			// gausa sama dengan karena dia 0 indexing 
			(*slice)[i] = 2 * i + 1
		} else{
			(*slice)[i] = 2 * (n - i) - 1

		}
	}
}

// Soal 8
func RepeatEvenPattern (slice *[]int){
	n := len(*slice)
	for i:= range (*slice){
		if (i < n/2){
			// gausa sama dengan karena dia 0 indexing 
			(*slice)[i] = 2 * (i+1)
			} else{
			(*slice)[i] = 2 * (n - i)

		}
	}
}

// Soal 9
func RepeatTriplePattern (slice *[]int){
	n := len(*slice)
	for i:= range (*slice){
		if (i < n/2){
			// gausa sama dengan karena dia 0 indexing 
			(*slice)[i] = 3 * (i+1)
			} else{
			(*slice)[i] = 3 * (n - i)

		}
	}
}

// Soal 11
func BuzzPattern(n int) {
	initVal := 3
	for i := range n {
		if (i % 2 == 0){
			fmt.Print("Buzz", "\t")
		} else {
			switch i {
			case 1:
				fmt.Print(1, "\t")
			default: 
				fmt.Print(initVal, "\t")
				initVal+=3

			}
		}
	}	
}

// Soal 12
func ModuloFourPattern(slice *[]int) {
	for i:= range (*slice) {
		(*slice)[i] = 2 * (i % 4) + 1
	}
}