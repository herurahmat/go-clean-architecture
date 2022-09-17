package book

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
	useCaseBook "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/book"
	"net/http"
)

type HandlerBook struct {
	Config  *config.Config
	Service useCaseBook.BookService
}

func NewBookHandler(c *config.Config, service useCaseBook.BookService) *HandlerBook {
	return &HandlerBook{
		Config:  c,
		Service: service,
	}
}

func (h *HandlerBook) GetBook(w http.ResponseWriter, r *http.Request) {

	data, err := h.Service.GetBook(r.Context())
	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get book",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}

func (h *HandlerBook) FindBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		config.Response(w, config.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusBadRequest)
	}

	data, err := h.Service.FindBookById(r.Context(), id)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get book",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}

func (h *HandlerBook) FindBookByName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	if name == "" {
		config.Response(w, config.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusBadRequest)
	}

	data, err := h.Service.FindBookByName(r.Context(), name)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get book",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}

func (h *HandlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var body book.BookModel
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusBadRequest)
	}

	author := author.AuthorModel{
		Id:   body.AuthorId,
		Name: "",
	}

	data, err := h.Service.CreateBook(r.Context(), author, body)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}

func (h *HandlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var body book.BookModel
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusBadRequest)
	}

	if body.Id == "" {
		config.Response(w, config.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusBadRequest)
	}

	author := author.AuthorModel{
		Id: body.AuthorId,
	}

	data, err := h.Service.UpdateBook(r.Context(), author, body)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}

func (h *HandlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {

	id := string(mux.Vars(r)["id"])

	config.Response(w, config.View{
		Status:       false,
		Code:         "400",
		Message:      "Bad Request",
		ErrorMessage: nil,
		Data:         nil,
		Pagination:   config.Pages{},
	}, http.StatusBadRequest)

	data, err := h.Service.DeleteBook(r.Context(), id)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   config.Pages{},
		}, http.StatusInternalServerError)
	}

	config.Response(w, config.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   config.Pages{},
	}, http.StatusOK)
}
