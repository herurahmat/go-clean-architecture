package author

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
)

type AuthorRepository interface {
	Get(ctx context.Context) (result []author.AuthorModel, err error)
	FindById(ctx context.Context, id string) (result author.AuthorModel, err error)
	FindByName(ctx context.Context, name string) (result author.AuthorModel, err error)
	Create(ctx context.Context, author *author.AuthorModel) (status bool, err error)
	Update(ctx context.Context, author *author.AuthorModel) (status bool, err error)
	Delete(ctx context.Context, id string) (status bool, err error)
}
