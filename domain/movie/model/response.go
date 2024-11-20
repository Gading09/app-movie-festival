package model

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
