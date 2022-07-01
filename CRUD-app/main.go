package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"ISBN"`
	Title    string    `json:"Title"`
	Director *Director `json:"Director"`
}
type Director struct {
	First_Name string `json:"First_Name"`
	LastName   string `json:"Last_Name"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "12312124", Title: "Once upon a time in Hollywood", Director: &Director{First_Name: "Quentin", LastName: "Tarantino"}})
	movies = append(movies, Movie{ID: "2", ISBN: "11111114", Title: "The Gentleman", Director: &Director{First_Name: "Guy", LastName: "Ritchie"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{ID}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{ID}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}
