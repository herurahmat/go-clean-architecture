package handler

import (
	"github.com/herurahmat/go-clean-architecture/internal/delivery/container"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/handler/author"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/handler/book"
)

type Handler struct {
	Author *author.HandlerAuthor
	Book   *book.HandlerBook
}

func New(container *container.Container) *Handler {
	author := author.NewAuthorHandler(container.Config, container.AuthorService)
	book := book.NewBookHandler(container.Config, container.BookService)
	return &Handler{
		Author: author,
		Book:   book,
	}
}
