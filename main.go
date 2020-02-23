package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-app/router"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		log.Println("middleware", r.URL)
		w.Header().Set("Content-Type", "application/json")
		// can be used for user token hadling.. CORS policy
		h.ServeHTTP(w, r)
	})
}

func main()  {
	fmt.Println("entry go")
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	router.RegisterHandlers(api)

	api.Use(Middleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
