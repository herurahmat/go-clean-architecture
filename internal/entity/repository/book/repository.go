package book

import (
	"context"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
)

type AuthorRepository interface {
	Get(ctx context.Context) (result []book.BookModel, err error)
	FindById(ctx context.Context, id string) (result book.BookModel, err error)
	FindByName(ctx context.Context, name string) (result book.BookModel, err error)
	Create(ctx context.Context, book *book.BookModel) (status bool, err error)
	Update(ctx context.Context, book *book.BookModel) (status bool, err error)
	Delete(ctx context.Context, id string) (status bool, err error)
}
