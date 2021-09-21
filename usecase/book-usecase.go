package usecase

import (
	"book-catalog/entity"
	"book-catalog/repository"
	"book-catalog/transport"
	"database/sql"
	"encoding/hex"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type BookUsecase interface {
	GetList() (*transport.GetList, *transport.ResponseError)
	Add(data transport.CreateBook) (*transport.GeneralResponse, *transport.ResponseError)
	Update(data transport.UpdateBook) (*transport.GeneralResponse, *transport.ResponseError)
	Delete(id string) (*transport.GeneralResponse, *transport.ResponseError)
	GetBook(id string) (*transport.GetBookResponse, *transport.ResponseError)
}
type bookUsecase struct {
	br repository.BookRepository
}

func NewBookUsecase(br repository.BookRepository) BookUsecase {
	return &bookUsecase{
		br: br,
	}
}

func (b *bookUsecase) GetList() (*transport.GetList, *transport.ResponseError) {
	result, err := b.br.GetList()
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}
	return &transport.GetList{
		Count:    len(result),
		ListBook: result,
	}, nil
}

func (b *bookUsecase) Add(data transport.CreateBook) (*transport.GeneralResponse, *transport.ResponseError) {
	// generate id: uuid
	id := uuid.NewV4()
	myuuid := hex.EncodeToString(id[:])
	createPayload := entity.Book{
		Id:      myuuid,
		Name:    data.Name,
		Creator: data.Creator,
	}

	err := b.br.Add(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	succes := &transport.GeneralResponse{
		Message: "succes",
	}
	return succes, nil
}

// update
func (b *bookUsecase) Update(data transport.UpdateBook) (*transport.GeneralResponse, *transport.ResponseError) {
	// make payload

	result, errBook := b.br.GetBook(data.Id)
	if errBook != nil {
		responseError := &transport.ResponseError{
			Message: "Not Found 404",
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	if data.Name == "" {
		data.Name = result.Name
	}
	if data.Creator == "" {
		data.Creator = result.Creator
	}

	createPayload := entity.Book{
		Id:      data.Id,
		Name:    data.Name,
		Creator: data.Creator,
	}

	err := b.br.Update(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}
	
	succes := &transport.GeneralResponse{
		Message: "succes",
	}
	return succes, nil

}

func (b *bookUsecase) Delete(id string) (*transport.GeneralResponse, *transport.ResponseError) {
	// find book id
	_, errBook := b.br.GetBook(id)
	if errBook != nil {
		responseError := &transport.ResponseError{
			Message: "Not Found 404",
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	err := b.br.Delete(id)
	if err != nil {
		responseError := &transport.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	return &transport.GeneralResponse{
		Message: "Delete Succes",
	}, nil
}

func (b *bookUsecase) GetBook(id string) (*transport.GetBookResponse, *transport.ResponseError) {
	result, err := b.br.GetBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			responseError := &transport.ResponseError{
				Message: "Not Found 404",
				Status:  http.StatusNotFound,
			}
			return nil, responseError
		}
	}

	return &transport.GetBookResponse{
		Data: *result,
	}, nil
}
