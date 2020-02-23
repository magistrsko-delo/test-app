package models

type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

func (book *Book) UpdateBook(book2 Book)  Book {
	book = &book2
	return *book
}
