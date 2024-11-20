package feature

import (
	"movie-festival/domain/movie/model"
	"movie-festival/domain/movie/repository"
	"time"

	"github.com/google/uuid"
)

type MovieFeature interface {
	CreateMovieFeature(request *model.ReqMovie) (err error)
}

type movieFeature struct {
	Repository repository.MovieRepository
}

func NewMovieFeature(repository repository.MovieRepository) MovieFeature {
	return &movieFeature{
		Repository: repository,
	}
}

func (feature movieFeature) CreateMovieFeature(request *model.ReqMovie) (err error) {
	genre := []model.Genre{}
	movieGenre := []model.MovieGenres{}
	for _, name := range request.Genres {
		id := uuid.New().String()
		data := feature.Repository.GetGenreByNameRepository(name)
		if data.Id == "" {
			genre = append(genre, model.Genre{
				Id:   id,
				Name: name,
			})
		} else {
			id = data.Id
		}
		movieGenre = append(movieGenre, model.MovieGenres{
			MovieId: request.Id,
			GenreId: id,
		})
	}

	artist := []model.Artist{}
	movieArtist := []model.MovieArtists{}
	for _, name := range request.Artists {
		id := uuid.New().String()
		data := feature.Repository.GetArtistByNameRepository(name)
		if data.Id == "" {
			artist = append(artist, model.Artist{
				Id:   id,
				Name: name,
			})
		} else {
			id = data.Id
		}
		movieArtist = append(movieArtist, model.MovieArtists{
			MovieId:  request.Id,
			ArtistId: id,
		})
	}

	return feature.Repository.CreateMovieRepository(
		model.Movie{
			Id:          request.Id,
			Title:       request.Title,
			Description: request.Description,
			WatchURL:    request.WatchUrl,
			Duration:    request.Duration,
			CreatedAt:   time.Now(),
		},
		genre,
		artist,
		movieGenre,
		movieArtist,
	)
}
