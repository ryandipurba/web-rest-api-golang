package main

import (
	"book-catalog/handler"
	"book-catalog/repository"
	"book-catalog/usecase"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewBookRepository()
	uCase := usecase.NewBookUsecase(repo)
	bookHandler := handler.NewBookHandler(uCase)

	r := mux.NewRouter()

	r.HandleFunc("/books", bookHandler.GetList).Methods("GET")
	r.HandleFunc("/books/{bookID}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books", bookHandler.AddBook).Methods("POST")
	r.HandleFunc("/books/{bookID}", bookHandler.UpdateBook).Methods("PATCH")
	r.HandleFunc("/books/{bookID}", bookHandler.DeleteBook).Methods("DELETE")

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}
