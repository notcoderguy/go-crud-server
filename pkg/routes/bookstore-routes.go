package routes

import (
	"github.com/gorilla/mux"
	"github.com/notcoderguy/go-crud-server/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/bookstore", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/bookstore/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/bookstore", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/bookstore/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/bookstore/{id}", controllers.UpdateBook).Methods("PUT")
}