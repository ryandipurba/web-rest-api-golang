package repository

import (
	"book-catalog/entity"
	"fmt"
)

type BookRepository interface {
	GetList() []entity.Book
	Add(payload entity.Book) bool
	Delete(id string) bool
	Update(payload entity.Book) bool
	GetBook(id string) entity.Book
}

type bookRepository struct {
	// book list
	book []entity.Book
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (b *bookRepository) GetList() []entity.Book {
	return b.book
}

func (b *bookRepository) Add(payload entity.Book) bool {
	b.book = append(b.book, payload)
	return true
}

// perulangan cuma untuk mencari index
func (b *bookRepository) Delete(id string) bool {
	for i, key := range b.book {
		if key.Id == id {
			b.book = append(b.book[:i], b.book[i+1:]...)
			return true
		}
	}
	// delete data in book list
	return false
}
func (b *bookRepository) Update(payload entity.Book) bool {
	// perulangan masih 2 kali
	for i, key := range b.book {
		fmt.Printf("perulangan %d \n", i+1)
		if key.Id == payload.Id {
			new := &b.book[i]
			if payload.Name != "" {
				new.Name = payload.Name
			}
			if payload.Creator != "" {
				new.Creator = payload.Creator
			}
			return true
		}
	}
	return false
}

func (b *bookRepository) GetBook(id string) entity.Book {
	var data entity.Book
	for _, key := range b.book {
		if key.Id == id {
			data = key
		}
	}
	// find data book in list
	// update data
	// update list
	return data
}
