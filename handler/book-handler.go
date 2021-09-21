package handler

import (
	"book-catalog/transport"
	"book-catalog/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type bookHandler struct {
	bu        usecase.BookUsecase
	validator *validator.Validate
}

func NewBookHandler(bu usecase.BookUsecase, validator *validator.Validate) *bookHandler {
	return &bookHandler{
		bu:        bu,
		validator: validator,
	}
}

// get list book
func (b *bookHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := b.bu.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// Add new book
func (b *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")
	var requestBook transport.CreateBook
	err := decoder.Decode(&requestBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responses := transport.ResponseError{
			Message: "error cuk",
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(responses)
		return
	}
	// validation
	errorValidation := b.validator.Struct(requestBook)
	if errorValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		dataResponse := transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(dataResponse)
		return
	}
	add, responseError := b.bu.Add(requestBook)
	if responseError != nil {
		w.WriteHeader(responseError.Status)
		json.NewEncoder(w).Encode(responseError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(add)
}

// // update book
func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["bookID"]
	var requestBook transport.UpdateBook
	err := decoder.Decode(&requestBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		responses := transport.ResponseError{
			Message: "error cuk",
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(responses)
		return
	}
	requestBook.Id = id
	// validation
	errorValidation := b.validator.Struct(requestBook)
	if errorValidation != nil {
		w.WriteHeader(http.StatusBadRequest)
		dataResponse := transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(dataResponse)
		return
	}
	update, responseError := b.bu.Update(requestBook)
	if responseError != nil {
		w.WriteHeader(responseError.Status)
		json.NewEncoder(w).Encode(responseError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(update)
}

// delete book
func (b *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["bookID"]
	delete, err := b.bu.Delete(id)
	if err != nil {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(delete)
}

func (b *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["bookID"]
	result, err := b.bu.GetBook(id)
	if err != nil {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
