package movie

import (
	"movie-festival/domain/movie/feature"
)

type MovieHandler interface {
}

type movieHandler struct {
	Feature feature.MovieFeature
}

func NewMovieHandler(feature feature.MovieFeature) MovieHandler {
	return &movieHandler{
		Feature: feature,
	}
}
