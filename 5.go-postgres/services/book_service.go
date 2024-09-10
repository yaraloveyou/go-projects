package services

import (
	"go-postgres/models"
	"go-postgres/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) GetById(id int) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) GetAll() ([]*models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) Update(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}
