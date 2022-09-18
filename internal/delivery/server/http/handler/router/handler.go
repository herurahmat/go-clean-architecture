package router

import (
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/handler"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/middleware"
	"net/http"
)

func NewRouter(r *mux.Router, c *config.Config, handler *handler.Handler) {

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Halo"))
	})

	authorRoute := r.PathPrefix("/author").Subrouter()
	authorRoute.HandleFunc("", handler.Author.GetAuthor).Methods(http.MethodGet)
	authorRoute.HandleFunc("/find-by-id/{id}", handler.Author.FindAuthorById).Methods(http.MethodGet)
	authorRoute.HandleFunc("", handler.Author.CreateAuthor).Methods(http.MethodPost)
	authorRoute.HandleFunc("/{id}", handler.Author.UpdateAuthor).Methods(http.MethodPut)
	authorRoute.HandleFunc("/{id}", handler.Author.DeleteAuthor).Methods(http.MethodDelete)
	authorRoute.Use(middleware.ApiKeyMiddleware(c))

	bookRoute := r.PathPrefix("/book").Subrouter()
	bookRoute.HandleFunc("", handler.Book.GetBook).Methods(http.MethodGet)
	bookRoute.HandleFunc("/find-by-id/{id}", handler.Book.FindBookById).Methods(http.MethodGet)
	bookRoute.HandleFunc("", handler.Book.CreateBook).Methods(http.MethodPost)
	bookRoute.HandleFunc("/{id}", handler.Book.UpdateBook).Methods(http.MethodPut)
	bookRoute.HandleFunc("/{id}", handler.Book.DeleteBook).Methods(http.MethodDelete)
	bookRoute.Use(middleware.ApiKeyMiddleware(c))
}
