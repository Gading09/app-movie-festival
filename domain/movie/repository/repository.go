package repository

import (
	"movie-festival/domain/movie/model"
	"movie-festival/helper/constant"
	e "movie-festival/helper/response/error"

	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovieRepository(
		payloadMovie model.Movie,
		payloadGenre []model.Genre,
		payloadArtist []model.Artist,
		movieGenre []model.MovieGenres,
		movieArtist []model.MovieArtists,
	) (err error)
	GetGenreByNameRepository(name string) (res model.Genre)
	GetArtistByNameRepository(name string) (res model.Artist)
}

type movieRepository struct {
	Database *gorm.DB
}

func NewMovieRepository(database *gorm.DB) MovieRepository {
	return &movieRepository{
		Database: database,
	}
}

func (repo movieRepository) CreateMovieRepository(
	payloadMovie model.Movie,
	payloadGenre []model.Genre,
	payloadArtist []model.Artist,
	movieGenre []model.MovieGenres,
	movieArtist []model.MovieArtists,
) (err error) {

	tx := repo.Database.Begin()

	if err := tx.Create(&payloadMovie).Error; err != nil {
		tx.Rollback()
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	if len(payloadGenre) != 0 {
		if err := tx.Create(&payloadGenre).Error; err != nil {
			tx.Rollback()
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return err
		}
	}

	if len(payloadArtist) != 0 {
		if err := tx.Create(&payloadArtist).Error; err != nil {
			tx.Rollback()
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return err
		}
	}

	if err := tx.Create(&movieGenre).Error; err != nil {
		tx.Rollback()
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	if err := tx.Create(&movieArtist).Error; err != nil {
		tx.Rollback()
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	return
}

func (repo movieRepository) GetGenreByNameRepository(name string) (res model.Genre) {
	repo.Database.Find(&res, "name = ?", name)
	return
}

func (repo movieRepository) GetArtistByNameRepository(name string) (res model.Artist) {
	repo.Database.Find(&res, "name = ?", name)
	return
}
