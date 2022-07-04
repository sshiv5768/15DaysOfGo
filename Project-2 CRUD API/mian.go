package main
import (
	"net/http"
	"log"
	"fmt"
	"math/rand"
	"encoding/json"
	"strconv"
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
}

func main(){
	movies := append(movies, 
		Movie{Id: 1, Title: "The Shawshank Redemption", Year: 1994, Director: "Frank Darabont"}, 
		Movie{Id: 2, Title: "The Godfather", Year: 1972, Director: "Francis Ford Coppola"}, 
		Movie{Id: 3, Title: "The Godfather: Part II", Year: 1974, Director: "Francis Ford Coppola"}, 
		Movie{Id: 4, Title: "The Dark Knight", Year: 2008, Director: "Christopher Nolan"},)
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).methods("GET")
	router.HandleFunc("/movies", addMovie).methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).methods("DELETE")

	fmt.Println("Server is starting on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
