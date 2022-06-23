// Day-3
// Constants and Output

package main
import ("fmt")

func main(){
	// Constant in grow are declared as "const constantName type = value"

	// You can avoid the type.
	
	const PI = 3.14
	
	const Hi = "Hello"
	const Bye = "Good Bye!"

	fmt.Print(PI, "\n") // Print() needs "/n" to leave a space behind
	fmt.Print(Hi,"\n", Bye, "\n")
	// In Println() you don't need "/n" to put newline at the end
	fmt.Printf("Hi's value is %v and its type is %t", Hi, Hi)
	// Printf() requires "%v" for value and %T for type of the variables.
}