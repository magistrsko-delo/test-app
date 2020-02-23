package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"test-app/models"
)

type BooksController struct {
	books []models.Book
}

func (BooksController *BooksController) GetBooks(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(BooksController.books)
	fmt.Println( (*(*BooksController).books[0].Author) )

	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode((*BooksController).books)
	fmt.Println("GETTING ALL books")
}

func (BooksController *BooksController) GetBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range BooksController.books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func (BooksController *BooksController) CreateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println( string(body)  )

	// _ = json.NewDecoder(r.Body).Decode(&book)
	json.Unmarshal(body, &book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID -- not save


	(*BooksController).books = append((*BooksController).books, book)
	json.NewEncoder(w).Encode(book)
}

func (BooksController *BooksController) UpdateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println( string(body)  )
	json.Unmarshal(body, &book)

	params := mux.Vars(r)
	for index, item := range BooksController.books {
		if item.ID == params["id"] {
			BooksController.books[index] = BooksController.books[index].UpdateBook(book)
			break
		}
	}

	json.NewEncoder(w).Encode((*BooksController).books)
}

func (BooksController *BooksController) DeleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range BooksController.books {
		if item.ID == params["id"] {
			BooksController.books = append(BooksController.books[:index], BooksController.books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode((*BooksController).books)
}


func InitController() *BooksController {
	return &BooksController{
		books: []models.Book {
			{
				ID:     "1",
				Isbn:   "45678",
				Title:  "Book ONe",
				Author: &models.Author{
					FirstName: "John",
					LastName:  "Foe",
				},
			},
			{
				ID:     "2",
				Isbn:   "34923740923479",
				Title:  "Book Two",
				Author: &models.Author{
					FirstName: "Steve",
					LastName:  "Smith",
				},
			},

		},
	}
}
