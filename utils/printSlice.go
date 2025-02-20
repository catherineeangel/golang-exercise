package utils

import "fmt"

func Print1DSlice (slice []int) {
	for i:= range len(slice){
		fmt.Print((slice)[i], " ")
	}
	fmt.Println()
}