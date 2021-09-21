package repository

import (
	"book-catalog/entity"
	"database/sql"
)

type BookRepository interface {
	GetList() ([]entity.Book, error)
	Add(payload entity.Book) error
	Delete(id string) error
	Update(payload entity.Book) error
	GetBook(id string) (*entity.Book, error)
}

type bookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}

func (b *bookRepository) GetList() ([]entity.Book, error) {

	rows, err := b.DB.Query("select * from tb_books")
	if err != nil {
		return nil, err
	}
	var books []entity.Book
	for rows.Next() {
		var res entity.Book
		_ = rows.Scan(&res.Id, &res.Name, &res.Creator)
		books = append(books, res)
	}
	return books, nil
}

func (b *bookRepository) Add(payload entity.Book) error {
	_, err := b.DB.Exec("INSERT INTO tb_books (id, name, creator) VALUES (?, ?, ?)", payload.Id, payload.Name, payload.Creator)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) Delete(id string) error {
	_, err := b.DB.Exec("DELETE FROM tb_books WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) Update(payload entity.Book) error {
	_, err := b.DB.Exec("UPDATE tb_books SET name = ?, creator = ? WHERE id = ?", payload.Name, payload.Creator, payload.Id)

	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) GetBook(id string) (*entity.Book, error) {
	sqlStatement := "SELECT * FROM tb_books WHERE id = ?"
	row := b.DB.QueryRow(sqlStatement, id)
	var book entity.Book
	err := row.Scan(&book.Id, &book.Name, &book.Creator)
	if err != nil {
		return nil, err
	}
	return &book, nil

}
