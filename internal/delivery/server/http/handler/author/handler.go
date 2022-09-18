package author

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	useCaseAuthor "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/author"
	"github.com/herurahmat/go-clean-architecture/internal/helper"
	"net/http"
)

type HandlerAuthor struct {
	Config  *config.Config
	Service useCaseAuthor.AuthorService
}

func NewAuthorHandler(c *config.Config, service useCaseAuthor.AuthorService) *HandlerAuthor {
	return &HandlerAuthor{
		Config:  c,
		Service: service,
	}
}

func (h *HandlerAuthor) GetAuthor(w http.ResponseWriter, r *http.Request) {

	data, err := h.Service.GetAuthor(r.Context())
	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get author",
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

func (h *HandlerAuthor) FindAuthorById(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Service.FindAuthorById(r.Context(), id)

	if err != nil {
		helper.ResponseHttp(w, helper.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get author",
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

func (h *HandlerAuthor) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body author.AuthorModel
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

	data, err := h.Service.CreateAuthor(r.Context(), body)

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

func (h *HandlerAuthor) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var body author.AuthorModel
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

	data, err := h.Service.UpdateAuthor(r.Context(), body.Id, body)

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

func (h *HandlerAuthor) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

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

	data, err := h.Service.DeleteAuthor(r.Context(), id)

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
