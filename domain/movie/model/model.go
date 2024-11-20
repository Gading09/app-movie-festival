package model

import "time"

type Movie struct {
	Id          string    `gorm:"primaryKey"`
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:255;not null"`
	WatchURL    string    `gorm:"size:255;not null"`
	Duration    string    `gorm:"size:10;not null"`
	ViewCount   int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type Genre struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;uniqueIndex:idx_genre_name;not null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type Artist struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;uniqueIndex:idx_artist_name;not null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
}

type MovieGenres struct {
	MovieId string
	GenreId string
}

type MovieArtists struct {
	MovieId  string
	ArtistId string
}

type Vote struct {
	UserId    string `gorm:"size:100;not null;uniqueIndex:idx_vote"`
	MovieId   string `gorm:"size:100;not null;uniqueIndex:idx_vote"`
	CreatedAt time.Time
}
