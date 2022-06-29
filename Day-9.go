// Day-9

package main
import ("fmt")

// Structure is similar to classes in other languages. Same as C language.
// It is a collection of data and functions of different types.

type student struct{
	name string
	class string
	subjects []string
	marks []int
}

func main(){
	var student student

	student.name = "Raj"
	student.class = "10"
	student.subjects = []string{"English", "Maths", "Science"}
	student.marks = []int{10, 20, 30}

	fmt.Println(student.marks[1])
	fmt.Println(student.name)

	

}