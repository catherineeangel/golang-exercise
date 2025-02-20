package logic02

import slice "github.com/catherineeangel/go-print-slice"

func Sample02 (){
	// init slice
	// n := 10
	slice2D := slice.CreateSlice(5)
	
	Soal01(&slice2D)
	slice.Print2DSlice(slice2D)

	Soal02(&slice2D)
	slice.Print2DSlice(slice2D)

	Soal03(&slice2D)
	slice.Print2DSlice(slice2D)

	sliceSoal6 := slice.CreateSlice(9)
	Soal06(&sliceSoal6)
	slice.Print2DSlice(sliceSoal6)

	sliceSoal6A := slice.CreateSlice(9)
	result := Soal06A(sliceSoal6A)
	slice.Print2DSlice(result)

	Soal07(&slice2D)
	slice.Print2DSlice(slice2D)

	slice2D = slice.CreateSlice(5)
	Soal08(&slice2D)
	slice.Print2DSlice(slice2D)

	Soal09(&slice2D)
	slice.Print2DSlice(slice2D)

	slice2D = slice.CreateSlice(5)
	Soal10(&slice2D)
	slice.Print2DSlice(slice2D)
	
	slice2D = slice.CreateSlice(5)
	Soal11(&slice2D)
	slice.Print2DSlice(slice2D)
	
	slice2D = slice.CreateSlice(6)
	Soal12(&slice2D)
	slice.Print2DSlice(slice2D)
	
	slice2D = slice.CreateSlice(5)
	Soal13(&slice2D)
	slice.Print2DSlice(slice2D)
}
