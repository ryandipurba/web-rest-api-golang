package handler

import (
	"book-catalog/transport"
	"book-catalog/usecase"
	"book-catalog/validation"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type bookHandler struct {
	bu usecase.BookUsecase
}

func NewBookHandler(bu usecase.BookUsecase) *bookHandler {
	return &bookHandler{
		bu: bu,
	}
}

// get list book
func (b *bookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts := b.bu.GetList()
	json.NewEncoder(w).Encode(posts)
}

// Add new book
func (b *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	w.Header().Set("Content-Type", "application/json")
	var requestBook transport.CreateBook
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &requestBook) //memasukan isi request body ke transport
	// validation
	err := validate.Struct(requestBook)
	if err != nil {
		errors := validation.FormatValidationError(err)
		dataResponse := transport.ValidateResponse{
			Message: errors,
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(dataResponse)
		return
	}
	add := b.bu.Add(requestBook)
	json.NewEncoder(w).Encode(add)
}

// update book
func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// validator buat global
	validate := validator.New()
	params := mux.Vars(r)
	id := params["bookID"]
	var requestBook transport.UpdateBook
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &requestBook)
	requestBook.Id = id
	// validation
	err := validate.Struct(requestBook)
	message := &transport.GeneralResponse{
		Message: "Field must be required",
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}
	requestBook.Id = id
	update := b.bu.Update(requestBook)
	json.NewEncoder(w).Encode(update)
}

// delete book
func (b *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["bookID"]
	deleteBook := b.bu.Delete(id)
	json.NewEncoder(w).Encode(deleteBook)
}

// get book
func (b *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["bookID"]
	getBook := b.bu.GetBook(id)
	json.NewEncoder(w).Encode(getBook)
}
