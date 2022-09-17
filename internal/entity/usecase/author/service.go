package author

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
)

type AuthorService interface {
	GetAuthor(ctx context.Context) (result []author.AuthorModel, err error)
	FindAuthorById(ctx context.Context, id string) (result author.AuthorModel, err error)
	FindAuthorByName(ctx context.Context, name string) (result author.AuthorModel, err error)
	CreateAuthor(ctx context.Context, dataEntity author.AuthorModel) (result author.AuthorModel, err error)
	UpdateAuthor(ctx context.Context, id string, dataEntity author.AuthorModel) (result author.AuthorModel, err error)
	DeleteAuthor(ctx context.Context, id string) (status bool, err error)
}
