//Day-4

package main
import ("fmt")

func main(){

	// Array Declaration and initialization
	// var arr_name = [length]type{value} - Length is specified.
	var arr1 = [3]int{2, 3, 4}

	// You can avoid specifying length compiler will take care of it.
	arr2 := []int{1, 2, 3, 4} // You can also declare array with ":=" operator

	fmt.Println(arr1)
	fmt.Println(arr2)

	// Length of array
	fmt.Println(len(arr2))

	// Changing array element with index number
	// You can also access elements with index number like fmt.Println(arr2[0])
	arr2[0] = 11
	fmt.Println(arr2)

}