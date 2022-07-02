package main

import (
	"net/http"
	"log"
)

func Hello(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h1 style='color: blue'>HellWorld</h1>"))
	
}

func main(){
	http.HandleFunc("/hello", Hello)

	log.Fatal(http.ListenAndServe(":5000", nil))
}