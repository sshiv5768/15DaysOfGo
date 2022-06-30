// Day-10

package main
import ("fmt")

func main(){
	// Maps in GO is a hash table. 
	// It is a data structure that stores key-value pairs.
	
	// Syntax: map[key-type]value-type{key: value, key: value, ...}

	var cars = map[string]string{"Maruti": "Swift", "Honda": "Civic"}
	fmt.Println(cars)

	// You can also use make() to create a map.

	var cars2 = make(map[string]string)

	// Add or update elements in the map.
	cars2["Ford"] = "Mustang"
	cars2["Lamborghini"] = "Huracan"
	fmt.Println(cars2["Ford"])

	// Use delete() to remove a key-value pair from a map.
	delete(cars2, "Ford")
	fmt.Println(cars2)

}
