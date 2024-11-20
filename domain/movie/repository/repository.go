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
	UpdateMovieRepository(
		payloadMovie model.Movie,
		payloadGenre []model.Genre,
		payloadArtist []model.Artist,
		movieGenre []model.MovieGenres,
		movieArtist []model.MovieArtists,
	) (err error)
	TopViewedMovieRepository() (res model.TopViewed, err error)
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

func (repo movieRepository) UpdateMovieRepository(
	payloadMovie model.Movie,
	payloadGenre []model.Genre,
	payloadArtist []model.Artist,
	movieGenre []model.MovieGenres,
	movieArtist []model.MovieArtists,
) (err error) {

	tx := repo.Database.Begin()

	if err := tx.Model(&model.Movie{}).Where("id = ?", payloadMovie.Id).Updates(payloadMovie).Error; err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	if len(payloadGenre) != 0 {
		if err := tx.Create(&payloadGenre).Error; err != nil {
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			tx.Rollback()
			return err
		}
	}

	if len(payloadArtist) != 0 {
		if err := tx.Create(&payloadArtist).Error; err != nil {
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			tx.Rollback()
			return err
		}
	}

	if len(movieGenre) != 0 {
		if err = tx.Where("movie_id = ?", payloadMovie.Id).Delete(model.MovieGenres{}).Error; err != nil {
			tx.Rollback()
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return
		}

		if err := tx.Create(&movieGenre).Error; err != nil {
			tx.Rollback()
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return err
		}
	}

	if len(movieGenre) != 0 {
		if err = tx.Where("movie_id = ?", payloadMovie.Id).Delete(model.MovieArtists{}).Error; err != nil {
			tx.Rollback()
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return
		}

		if err := tx.Create(&movieArtist).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return err
	}

	return
}

func (repo movieRepository) TopViewedMovieRepository() (res model.TopViewed, err error) {
	var mostVotedMovie model.MostVotedMovie
	var mostViewedGenre model.MostViewedGenre

	movieResult := repo.Database.Table("votes").
		Select("movie_id, movies.title, count(*) as voted").
		Group("movie_id").
		Order("voted DESC").
		Limit(1).
		Joins("JOIN movies ON movies.id = votes.movie_id").
		Scan(&mostVotedMovie)

	if movieResult.Error != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	genreResult := repo.Database.Table("movie_genres").
		Select("genres.id as genre_id, genres.name as name, SUM(movies.view_count) as view_count").
		Joins("JOIN genres ON genres.id = movie_genres.genre_id").
		Joins("JOIN movies ON movies.id = movie_genres.movie_id").
		Group("genres.id, genres.name").
		Order("view_count DESC").
		Limit(1).
		Scan(&mostViewedGenre)

	if genreResult.Error != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	res = model.TopViewed{
		Movie: mostVotedMovie,
		Genre: mostViewedGenre,
	}
	return
}
