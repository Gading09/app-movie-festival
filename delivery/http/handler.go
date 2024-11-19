package http

import (
	"movie-festival/delivery/container"
)

type Handler struct {
}

func SetupHandler(container container.Container) Handler {
	return Handler{}
}
