// Day-8

package main
import "fmt"

// In go functions are same as in other languages.
func myFunction(){
	fmt.Println("myFunction Just Executed!")
}
// Function with parameters and arguments
// name1 and name2 are parameters and Raj and Mohan are arguments.

func names(name1 string, name2 string){
	fmt.Println("Calling", name1, "and", name2)
}

// Function with return value
func add(x int, y int) int /*return type*/{
	return x + y
}

func main(){
	myFunction()
	names("Raj", "Mohan")
	fmt.Println(add(10, 20))
	fmt.Println("main function just executed!")

}