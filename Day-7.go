// Day-7

package main
import ("fmt")

func main(){

	// Go switch statement
	// It is same as it is in other languages.
	var day = 2
	switch day{
	case 1:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Wednesday")	
	default:
		fmt.Println("Invalid day")
	}	

	// Go for loop
	// There is no do-while loop in GO. We always use for loop.
	
	// Type-1 Normal for loop
	// Syntax: for statement1; condition; increment/decrement{}
	for i:= 0; i < 5; i++{
		fmt.Println(i)
	}

	// Type-2 Range for loop
	// Mainly used for arrays and slices.
	// Syntax: for index, value := range array/slice{}
	var arr = []int{2, 4, 6, 8}
	for index, value := range arr{
		fmt.Println(index, value)
	}

	// if you want to omit index, you can use _ instead of index. Same for value also.

	// Type-3 While kinda loop
	// Syntax: for condition{}
	var j = 0
	for j < 5{
		fmt.Println(j)
		j++
	}



}