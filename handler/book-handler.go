package handler

import (
	"book-catalog/entity"
	"book-catalog/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type BookHandler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
}
type bookHandler struct {
	bu usecase.BookUsecase
}

func NewBooHandler(bu usecase.BookUsecase) BookHandler {
	return &bookHandler{
		bu: bu,
	}
}

func (b *bookHandler) GetList(w http.ResponseWriter, _ *http.Request) {
	object := b.bu.GetList()
	json.NewEncoder(w).Encode(object)
}

func (b *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t entity.Book
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		b.bu.AddBook(t.Id, t.Name, t.Creator)
		object := b.bu.GetList()
		log.Println(object)
		json.NewEncoder(w).Encode(t)
	}
}
