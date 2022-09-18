package http

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/container"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/handler"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server/http/handler/router"
	"log"
	"net/http"
)

func NewHttp(container *container.Container) {
	r := mux.NewRouter()
	router.NewRouter(r, container.Config, handler.New(container))

	fmt.Print("Running ", container.Config.Server.Ip, " ", container.Config.Server.Port)
	log.Println(http.ListenAndServe(container.Config.Server.Ip+":"+container.Config.Server.Port, handlers.CompressHandler(r)))
}
