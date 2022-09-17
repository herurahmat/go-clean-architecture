package author

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/config"
	"github.com/herurahmat/go-clean-architecture/internal/entity/author"
	useCaseAuthor "github.com/herurahmat/go-clean-architecture/internal/entity/usecase/author"
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
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get author",
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

func (h *HandlerAuthor) FindAuthorById(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Service.FindAuthorById(r.Context(), id)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get author",
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

func (h *HandlerAuthor) FindAuthorByName(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Service.FindAuthorByName(r.Context(), name)

	if err != nil {
		config.Response(w, config.View{
			Status:       false,
			Code:         "500",
			Message:      "Failed get author",
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

func (h *HandlerAuthor) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var body author.AuthorModel
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

	data, err := h.Service.CreateAuthor(r.Context(), body)

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

func (h *HandlerAuthor) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var body author.AuthorModel
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

	body.Id = string(mux.Vars(r)["id"])
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

	data, err := h.Service.UpdateAuthor(r.Context(), body.Id, body)

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

func (h *HandlerAuthor) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	id := string(mux.Vars(r)["id"])

	config.Response(w, config.View{
		Status:       false,
		Code:         "400",
		Message:      "Bad Request",
		ErrorMessage: nil,
		Data:         nil,
		Pagination:   config.Pages{},
	}, http.StatusBadRequest)

	data, err := h.Service.DeleteAuthor(r.Context(), id)

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
