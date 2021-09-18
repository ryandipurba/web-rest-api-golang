package usecase

import (
	"book-catalog/entity"
	"book-catalog/repository"
)

type BookUsecase interface {
	AddBook(id int, name string, creator string) bool
	GetList() []entity.Book
}
type bookUsecase struct {
	br repository.BookRepository
}

func NewBookUsecase(br repository.BookRepository) BookUsecase {
	return &bookUsecase{
		br: br,
	}
}

func (b *bookUsecase) AddBook(id int, name string, creator string) bool {
	p := entity.Book{
		Id:      id,
		Name:    name,
		Creator: creator,
	}

	b.br.AddBook(p)
	return true

}

func (b *bookUsecase) GetList() []entity.Book {
	return b.br.GetList()
}
