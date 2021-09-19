package usecase

import (
	"book-catalog/entity"
	"book-catalog/repository"
	"book-catalog/transport"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

type BookUsecase interface {
	GetList() transport.GetList
	Add(data transport.CreateBook) transport.GeneralResponse
	Update(data transport.UpdateBook) transport.GeneralResponse
	Delete(id string) transport.GeneralResponse
	GetBook(id string) transport.GetBook
}
type bookUsecase struct {
	br repository.BookRepository
}

func NewBookUsecase(br repository.BookRepository) BookUsecase {
	return &bookUsecase{
		br: br,
	}
}

func (b *bookUsecase) GetList() transport.GetList {
	data := b.br.GetList()
	return transport.GetList{
		Count:    len(data),
		ListBook: data,
	}
}

func (b *bookUsecase) Add(data transport.CreateBook) transport.GeneralResponse {
	// generate id: uuid
	id := uuid.NewV4()
	myuuid := hex.EncodeToString(id[:])
	createPayload := entity.Book{
		Id:      myuuid,
		Name:    data.Name,
		Creator: data.Creator,
	}

	// save data to repostory
	// make response
	b.br.Add(createPayload)
	return transport.GeneralResponse{
		Message: "berhasil",
	}
}

func (b *bookUsecase) Update(data transport.UpdateBook) transport.GeneralResponse {
	// make payload
	createPayload := entity.Book{
		Id:      data.Id,
		Name:    data.Name,
		Creator: data.Creator,
	}

	if !b.br.Update(createPayload) {
		return transport.GeneralResponse{
			Message: "Id Not Found",
		}
	}
	// update data in repository
	return transport.GeneralResponse{
		Message: "Succes",
	}
}

func (b *bookUsecase) Delete(id string) transport.GeneralResponse {
	// delete data in repostiory
	// if !b.br.Delete(id) {
	// 	return transport.GeneralResponse{
	// 		Message: "Id Not Found",
	// 	}
	// }

	b.br.Delete(id)

	return transport.GeneralResponse{
		Message: "Succes",
	}
}

func (b *bookUsecase) GetBook(id string) transport.GetBook {
	// delete data in repostiory
	data := b.br.GetBook(id)
	return transport.GetBook{
		Data: data,
	}
}
