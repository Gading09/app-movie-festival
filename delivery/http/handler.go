package http

import (
	"movie-festival/delivery/container"
	movie "movie-festival/domain/movie"
	user "movie-festival/domain/user"
)

type Handler struct {
	MovieHandler movie.MovieHandler
	UserHandler  user.UserHandler
}

func SetupHandler(container container.Container) Handler {
	return Handler{
		MovieHandler: movie.NewMovieHandler(container.MovieFeature),
		UserHandler:  user.NewUserHandler(container.UserFeature),
	}
}
