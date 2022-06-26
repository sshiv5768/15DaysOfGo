// Day-6
package main
import ("fmt")

func main(){
	// Go conditions
	var x = 10
	var y = 20
	// Unlike other languages, in GO, else and else if block should not be in newline.

	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("Both are equal")
	}
	

}
