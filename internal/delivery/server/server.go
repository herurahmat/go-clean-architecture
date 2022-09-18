package server

import (
	"github.com/herurahmat/go-clean-architecture/internal/delivery/container"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http"
)

func New(container *container.Container) {
	http.NewHttp(container)
}
