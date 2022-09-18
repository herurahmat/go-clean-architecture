package main

import (
	"github.com/herurahmat/go-clean-architecture/internal/delivery/container"
	"github.com/herurahmat/go-clean-architecture/internal/delivery/server"
)

func main() {
	server.New(container.New())
}
