package book

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/entity/book"
	useCaseBook "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/book"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
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
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get book",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseHttp(w, helper.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   helper.Pages{},
	}, http.StatusOK)
	return
}

func (h *HandlerBook) FindBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.FindBookById(r.Context(), id)

	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get book",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseHttp(w, helper.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   helper.Pages{},
	}, http.StatusOK)
	return
}

func (h *HandlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var body book.BookModel
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.CreateBook(r.Context(), body)

	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseHttp(w, helper.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   helper.Pages{},
	}, http.StatusOK)
	return
}

func (h *HandlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var body book.BookModel
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: err,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusBadRequest)
		return
	}
	body.Id = string(mux.Vars(r)["id"])
	if body.Id == "" {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.UpdateBook(r.Context(), body)

	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseHttp(w, helper.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   helper.Pages{},
	}, http.StatusOK)
	return
}

func (h *HandlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {

	id := string(mux.Vars(r)["id"])

	if id == "" {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "400",
			Message:      "Bad Request",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusBadRequest)
		return
	}

	data, err := h.Service.DeleteBook(r.Context(), id)

	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Internal server exception",
			ErrorMessage: nil,
			Data:         nil,
			Pagination:   helper.Pages{},
		}, http.StatusInternalServerError)
		return
	}

	helper.ResponseHttp(w, helper.View{
		Status:       true,
		Code:         "200",
		Message:      "Success",
		ErrorMessage: nil,
		Data:         data,
		Pagination:   helper.Pages{},
	}, http.StatusOK)
	return
}
