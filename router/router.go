package router

import (
	"github.com/gorilla/mux"
	"test-app/controllers"
)

type CustomRouter struct {
	R *mux.Router
}

func (customRouter *CustomRouter) RegisterHandlers(/*r *mux.Router*/) {
	bookController := controllers.InitController()

	(*customRouter).R.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	(*customRouter).R.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	(*customRouter).R.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	(*customRouter).R.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	(*customRouter).R.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

}