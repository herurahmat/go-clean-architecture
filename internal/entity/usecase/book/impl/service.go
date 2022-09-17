package impl

import (
	"context"
	"database/sql"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	model "github.com/herurahmat/go-clean-architecture/internal/entity/book"
	repository "github.com/herurahmat/go-clean-architecture/internal/entity/repository/book"
	"log"
)

type service struct {
	repo   repository.BookRepository
	db     *sql.DB
	config *config.Config
}

func NewServiceBook(repo repository.BookRepository, db *sql.DB, config *config.Config) *service {
	return &service{
		repo:   repo,
		db:     db,
		config: config,
	}
}

func (s *service) GetBook(ctx context.Context) (result []model.BookModel, err error) {
	data, err := s.repo.Get(ctx)

	if err != nil {
		log.Println("error", err)
		return nil, nil
	}

	return data, nil
}

func (s *service) FindBookById(ctx context.Context, id string) (result model.BookModel, err error) {

	if id == "" {
		log.Println("Id empty")
		return model.BookModel{}, nil
	}

	data, err := s.repo.FindById(ctx, id)

	if err != nil {
		log.Println("id ", id, " name", err)
		return model.BookModel{}, nil
	}

	return data, nil
}

func (s *service) FindBookByName(ctx context.Context, name string) (result model.BookModel, err error) {
	if name == "" {
		log.Println("name empty")
		return model.BookModel{}, nil
	}

	data, err := s.repo.FindByName(ctx, name)

	if err != nil {
		log.Println("name ", name, " error ", err)
		return model.BookModel{}, nil
	}

	return data, nil
}

func (s *service) CreateBook(ctx context.Context, author author.AuthorModel, book model.BookModel) (result model.BookModel, err error) {

	if book.GetBookTitle() == "" {
		log.Println("book empty")
	}

	data, err := s.repo.Create(ctx, author, book)

	if err != nil {
		log.Println("error:", err)
		return model.BookModel{}, nil
	}

	return data, nil
}

func (s *service) UpdateBook(ctx context.Context, author author.AuthorModel, book model.BookModel) (result model.BookModel, err error) {
	if book.GetBookTitle() == "" {
		log.Println("book empty")
	}

	data, err := s.repo.Update(ctx, author, book)

	if err != nil {
		log.Println("error:", err)
		return model.BookModel{}, nil
	}

	return data, nil
}

func (s *service) DeleteBook(ctx context.Context, id string) (status bool, err error) {
	if id == "" {
		log.Println("Id empty")
		return false, nil
	}

	status, err = s.repo.Delete(ctx, id)

	if err != nil {
		log.Println("error:", err)
		return false, nil
	}

	return status, nil
}
