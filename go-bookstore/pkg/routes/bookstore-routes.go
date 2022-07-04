package routes

import (
	"github.com/AlexS778/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoresRoutes = func(router mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
