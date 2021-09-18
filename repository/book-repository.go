package repository

import (
	"book-catalog/entity"
)

type BookRepository interface {
	AddBook(payload entity.Book) bool
	GetList() []entity.Book
}

type bookRepository struct {
	book []entity.Book
}

func NewBajuRepository() BookRepository {
	return &bookRepository{}
}

func (b *bookRepository) AddBook(payload entity.Book) bool {
	b.book = append(b.book, payload)
	return true
}

func (b *bookRepository) GetList() []entity.Book {
	b.book = []entity.Book{
		{
			Id:      1,
			Name:    "Study go",
			Creator: "ryan",
		},
		{
			Id:      2,
			Name:    "Study js",
			Creator: "anto",
		},
		{
			Id:      3,
			Name:    "Study html",
			Creator: "jaka",
		},
	}
	return b.book
}
