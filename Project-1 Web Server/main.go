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
}