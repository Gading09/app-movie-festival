package feature

import (
	"movie-festival/domain/movie/repository"
)

type MovieFeature interface {
}

type movieFeature struct {
	Repository repository.MovieRepository
}

func NewMovieFeature(repository repository.MovieRepository) MovieFeature {
	return &movieFeature{
		Repository: repository,
	}
}
