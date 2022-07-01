package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	// Creating a new file server
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Creating a listenerAndServe at port 8080
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}