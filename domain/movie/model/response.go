package model

import "time"

type MostVotedMovie struct {
	MovieId string `json:"movie_id"`
	Title   string `json:"title"`
	Voted   int    `json:"voted"`
}

type MostViewedGenre struct {
	GenreId   string `json:"genre_id"`
	Name      string `json:"name"`
	ViewCount int    `json:"view_count"`
}

type TopViewed struct {
	Movie MostVotedMovie  `json:"most_voted_movie"`
	Genre MostViewedGenre `json:"most_viewed_genre"`
}

type GetListMovie struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Genre       []string  `json:"genres"`
	Artist      []string  `json:"artits"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Pagination struct {
	Limit     int `json:"limit_per_page"`
	Page      int `json:"current_page"`
	TotalPage int `json:"total_page"`
	TotalRows int `json:"total_rows"`
}

type ResGetListMovie struct {
	Pagination Pagination     `json:"pagination"`
	Data       []GetListMovie `json:"data"`
}

type ResGetListMovieBySearch struct {
	Data []GetListMovie `json:"data"`
}

type WatchMovie struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	WatchUrl  string `json:"watch_url"`
	ViewCount int    `json:"view_count"`
}
