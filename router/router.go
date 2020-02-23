package router

import (
	"github.com/gorilla/mux"
	"test-app/controllers"
)

func RegisterHandlers(r *mux.Router) {
	bookController := controllers.InitController()

	r.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

}