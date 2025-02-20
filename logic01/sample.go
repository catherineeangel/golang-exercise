package logic01

import "github.com/catherineeangel/golang-exercise/utils"

func Sample01 () {
	oddSlide := OddPattern(7)
	evenSlide := EvenPattern(7)
	utils.Print1DSlice(oddSlide)
	utils.Print1DSlice(evenSlide)
}