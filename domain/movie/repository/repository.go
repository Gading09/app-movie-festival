package repository

import (
	"gorm.io/gorm"
)

type MovieRepository interface {
}

type movieRepository struct {
	Database *gorm.DB
}

func NewMovieRepository(database *gorm.DB) MovieRepository {
	return &movieRepository{
		Database: database,
	}
}
