package logic01

// import "github.com/catherineeangel/golang-exercise/utils"
import (
	"fmt"

	slice "github.com/catherineeangel/go-print-slice"
)

func Sample01 () {
	array := generateSlice(10)

	fmt.Println("Odd Slice")
	OddPattern(&array)
	slice.Print1DSlice(array)
	fmt.Println("Even Slice")
	EvenPattern(&array)
	slice.Print1DSlice(array)
	fmt.Println("Modulo Four Pattern")
	ModuloFourPattern(&array)
	slice.Print1DSlice(array)
	fmt.Println("Reverse Odd Pattern")
	ReverseOddPattern(&array)
	slice.Print1DSlice(array)
	
	fmt.Println("Repeating Odd Pattern")
	RepeatOddPattern(&array)
	slice.Print1DSlice(array)
	fmt.Println("Repeating Even Pattern")
	RepeatEvenPattern(&array)
	slice.Print1DSlice(array)
	fmt.Println("Repeating Triple Pattern")
	RepeatTriplePattern(&array)
	slice.Print1DSlice(array)

	
}

func generateSlice(n int) []int {
	return make([]int, n)
}