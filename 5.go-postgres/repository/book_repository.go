package repository

import (
	"database/sql"
	"go-postgres/models"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	query := "INSERT INTO books (title, author, isbn) VALUES ($1, $2, $3) RETURNING id"
	err := r.DB.QueryRow(query, book.Title, book.Author, book.ISBN).Scan(&book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) GetByID(id int) (*models.Book, error) {
	book := &models.Book{}
	query := "SELECT id, title, author, isbn FROM books WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) GetAll() ([]*models.Book, error) {
	var books []*models.Book
	rows, err := r.DB.Query("SELECT id, title, author, isbn FROM books")
	if err != nil {
		return books, err
	}
	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN)
		if err != nil {
			return books, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	query := "UPDATE books SET title = $1, author = $2, isbn = $3 WHERE id = $4"
	_, err := r.DB.Exec(query, book.Title, book.Author, book.ISBN, book.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) Delete(id int) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
