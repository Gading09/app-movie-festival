package repository

import (
	"movie-festival/domain/movie/model"
	"movie-festival/helper/constant"
	e "movie-festival/helper/response/error"
	"strings"

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
	GetTotalDataRepository() (count int64, err error)
	GetListMovieRepository(limit, offset int) (res []model.GetListMovie, err error)
	GetListMovieBySearchRepository(search string) (res []model.GetListMovie, err error)
	GetMovieByIdRepository(id string) (res model.Movie, err error)
	IncViewMovieRepository(id string) (err error)
	VoteMovieRepository(payload model.Vote) (err error)
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

func (repo movieRepository) GetTotalDataRepository() (count int64, err error) {
	if err = repo.Database.Model(&model.Movie{}).Count(&count).Error; err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	return
}

func (repo movieRepository) GetListMovieRepository(limit, offset int) (res []model.GetListMovie, err error) {
	rows, err := repo.Database.Table("movies").
		Select(`movies.id, movies.title, movies.created_at, 
                GROUP_CONCAT(DISTINCT genres.name) AS genre_names, 
                GROUP_CONCAT(DISTINCT artists.name) AS artist_names`).
		Joins("JOIN movie_genres ON movies.id = movie_genres.movie_id").
		Joins("JOIN genres ON movie_genres.genre_id = genres.id").
		Joins("JOIN movie_artists ON movies.id = movie_artists.movie_id").
		Joins("JOIN artists ON movie_artists.artist_id = artists.id").
		Group("movies.id").
		Limit(limit).
		Offset(offset).
		Rows()
	if err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var movie model.GetListMovie
		var genreNames string
		var artistNames string
		if err = rows.Scan(&movie.Id, &movie.Title, &movie.CreatedAt, &genreNames, &artistNames); err != nil {
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return
		}

		movie.Genre = strings.Split(genreNames, ",")
		movie.Artist = strings.Split(artistNames, ",")
		res = append(res, movie)
	}

	return
}

func (repo movieRepository) GetListMovieBySearchRepository(search string) (res []model.GetListMovie, err error) {
	query := repo.Database.Table("movies").
		Select(`movies.id, movies.title, movies.created_at, movies.description, 
                GROUP_CONCAT(distinct genres.name) AS genre_names, 
                GROUP_CONCAT(distinct artists.name) AS artist_names`).
		Joins("JOIN movie_genres ON movies.id = movie_genres.movie_id").
		Joins("JOIN genres ON movie_genres.genre_id = genres.id").
		Joins("JOIN movie_artists ON movies.id = movie_artists.movie_id").
		Joins("JOIN artists ON movie_artists.artist_id = artists.id").
		Group("movies.id")

	if search != "" {
		query = query.Having("movies.title LIKE ? OR movies.description LIKE ? OR artist_names LIKE ? OR genre_names LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	rows, err := query.Rows()
	if err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return

	}

	for rows.Next() {
		var movie model.GetListMovie
		var genreNames string
		var artistNames string
		if err = rows.Scan(&movie.Id, &movie.Title, &movie.CreatedAt, &movie.Description, &genreNames, &artistNames); err != nil {
			err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
			return
		}

		movie.Genre = strings.Split(genreNames, ",")
		movie.Artist = strings.Split(artistNames, ",")
		res = append(res, movie)
	}
	defer rows.Close()
	return
}

func (repo movieRepository) GetMovieByIdRepository(id string) (res model.Movie, err error) {
	if err = repo.Database.First(&res, "id = ?", id).Error; err != nil {
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}
	return
}

func (repo movieRepository) IncViewMovieRepository(id string) (err error) {
	if err = repo.Database.Model(&model.Movie{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + ?", 1)).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = e.New(constant.StatusNotFound, constant.ErrNotFound, err)
			return
		}
		return err
	}
	return
}

func (repo movieRepository) VoteMovieRepository(payload model.Vote) (err error) {
	if err = repo.Database.Create(&payload).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			err = e.New(constant.StatusInternalServerError, constant.ErrAlreadyExist, err)
			return
		}
		err = e.New(constant.StatusInternalServerError, constant.ErrDatabase, err)
		return
	}

	return
}
