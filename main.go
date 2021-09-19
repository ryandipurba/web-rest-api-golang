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

	/**
	method : GET
	request: name
	response : {
		count : number
		listBook : {
			id : "xxx" -> format uuid
			name : "new book"
			creator: "vani"
		}
	}

	*/
	r.HandleFunc("/books", bookHandler.GetList).Methods("GET")
	r.HandleFunc("/books/{bookID}", bookHandler.GetBook).Methods("GET")
	/**
	method : POST
	request: {
		name: "new book", -> required
		creator: "ryand", ->required
	}
	response : {
		message : "berhasil"
	}
	*/
	r.HandleFunc("/books", bookHandler.Add).Methods("POST")
	/**
	method : PATCH
	request: {
		bookID: "xxx", -> required
		name: "new book",-> optional
		creator: "ryand",-> optional
	}
	response : {
		message : "berhasil"
	}
	*/
	r.HandleFunc("/books/{bookID}", bookHandler.UpdateBook).Methods("PATCH")
	/**
	method : Delete
	response : {
		message : "berhasil"
	}
	*/
	r.HandleFunc("/books/{bookID}", bookHandler.DeleteBook).Methods("DELETE")

	fmt.Println("Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}
