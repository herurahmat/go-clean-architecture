package impl

import (
	"context"
	"database/sql"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	model "github.com/herurahmat/go-clean-architecture/internal/entity/author"
	repository "github.com/herurahmat/go-clean-architecture/internal/entity/repository/author"
	"log"
)

type service struct {
	repo   repository.AuthorRepository
	db     *sql.DB
	config *config.Config
}

func NewServiceAuthor(repo repository.AuthorRepository, db *sql.DB, config *config.Config) *service {
	return &service{
		repo:   repo,
		db:     db,
		config: config,
	}
}

func (s *service) GetAuthor(ctx context.Context) (result []model.AuthorModel, err error) {
	data, err := s.repo.Get(ctx)

	if err != nil {
		log.Println("error", err)
		return nil, nil
	}

	return data, nil
}

func (s *service) FindAuthorById(ctx context.Context, id string) (result model.AuthorModel, err error) {

	if id == "" {
		log.Println("Id empty")
		return model.AuthorModel{}, nil
	}

	data, err := s.repo.FindById(ctx, id)

	if err != nil {
		log.Println("id ", id, " name", err)
		return model.AuthorModel{}, nil
	}

	return data, nil
}

func (s *service) CreateAuthor(ctx context.Context, author model.AuthorModel) (result model.AuthorModel, err error) {

	if author.GetAuthorName() == "" {
		log.Println("author empty")
	}

	data, err := s.repo.Create(ctx, author)

	if err != nil {
		log.Println("error:", err)
		return model.AuthorModel{}, nil
	}

	return data, nil
}

func (s *service) UpdateAuthor(ctx context.Context, id string, author model.AuthorModel) (result model.AuthorModel, err error) {
	if author.GetAuthorId() == "" {
		log.Println("author empty")
	}

	data, err := s.repo.Update(ctx, id, author)

	if err != nil {
		log.Println("error:", err)
		return model.AuthorModel{}, nil
	}

	return data, nil
}

func (s *service) DeleteAuthor(ctx context.Context, id string) (status bool, err error) {
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
