package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoresRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9920", r))
}
