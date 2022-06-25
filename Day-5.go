// Day-5 
package main
import ("fmt")

func main(){
	// Slices- Powerful tool in GOlang
	// Works same as array but it's length is fixed.

	// Declaration
	/* 
	slice_name := []datatype{value}
	
	Or
	slice_name := array[start:end]
	
	Or
	slice_name = make([]type, length, capacity)

	*/

	slice1 := []int{1,2,3}
	fmt.Println(slice1, len(slice1), cap(slice1))

	var arr1 = []int{1, 2, 13, 14, 122}
	var slice2 = arr1[1:4]
	fmt.Println(slice2, len(slice2), cap(slice2))

	slice3 := make([]int, 3, 3)
	fmt.Println(slice3, len(slice3), cap(slice3))
	

}