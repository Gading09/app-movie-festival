package feature

import (
	"movie-festival/domain/movie/model"
	"movie-festival/domain/movie/repository"
	"movie-festival/helper"
	"time"

	"github.com/google/uuid"
)

type MovieFeature interface {
	CreateMovieFeature(request *model.ReqMovie) (err error)
	UpdateMovieFeature(request *model.ReqUpdateMovie) (err error)
	TopViewedMovieFeature() (res model.TopViewed, err error)
	GetListMovieFeature(request *model.ReqGetListMovie) (res model.ResGetListMovie, err error)
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

func (feature movieFeature) UpdateMovieFeature(request *model.ReqUpdateMovie) (err error) {
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
	return feature.Repository.UpdateMovieRepository(
		model.Movie{
			Id:          request.Id,
			Title:       request.Title,
			Description: request.Description,
			WatchURL:    request.WatchUrl,
			Duration:    request.Duration,
		},
		genre,
		artist,
		movieGenre,
		movieArtist,
	)
}

func (feature movieFeature) TopViewedMovieFeature() (res model.TopViewed, err error) {
	return feature.Repository.TopViewedMovieRepository()
}

func (feature movieFeature) GetListMovieFeature(request *model.ReqGetListMovie) (res model.ResGetListMovie, err error) {
	totalData, err := feature.Repository.GetTotalDataRepository()
	if err != nil {
		return
	}

	offset, totalPage := helper.GetPaginations(int(totalData), request.Limit, request.Page)
	items, err := feature.Repository.GetListMovieRepository(request.Limit, offset)
	if err != nil {
		return
	}
	return model.ResGetListMovie{
		Pagination: model.Pagination{
			Limit:     request.Limit,
			Page:      request.Page,
			TotalPage: totalPage,
			TotalRows: int(totalData),
		},
		Data: items,
	}, nil
}
