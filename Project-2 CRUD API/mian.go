package main
import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Movie struct{
	Id int `json:"id"`
	Title string `json: "title"`
	Year int `json: "year"`
	Director string `json: "director"`

}
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
/*
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.Id == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		}
	}
	json.NewEncoder(w).Encode(movies)	

*/
func main(){
	movies = append(movies, 
		Movie{Id: 1, Title: "The Shawshank Redemption", Year: 1994, Director: "Frank Darabont"})

	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	


	fmt.Println("Server is starting on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
