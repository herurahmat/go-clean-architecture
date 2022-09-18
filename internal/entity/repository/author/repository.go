package author

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
)

type AuthorRepository interface {
	Get(ctx context.Context) (result []author.AuthorModel, err error)
	FindById(ctx context.Context, id string) (result author.AuthorModel, err error)
	Create(ctx context.Context, dataEntity author.AuthorModel) (author.AuthorModel, error)
	Update(ctx context.Context, id string, dataEntity author.AuthorModel) (author.AuthorModel, error)
	Delete(ctx context.Context, id string) (status bool, err error)
}
