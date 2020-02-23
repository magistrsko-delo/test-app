package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main()  {
	fmt.Println("entry go")
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
