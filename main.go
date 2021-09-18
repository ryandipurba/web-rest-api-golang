package main

import (
	"book-catalog/handler"
	"book-catalog/repository"
	"book-catalog/usecase"
	"net/http"
)

func main() {

	repo := repository.NewBajuRepository()
	uCase := usecase.NewBookUsecase(repo)
	bookHandler := handler.NewBooHandler(uCase)

	http.HandleFunc("/books", bookHandler.GetList)
	http.HandleFunc("/book/post", bookHandler.AddBook)
	http.ListenAndServe(":8080", nil)
}
