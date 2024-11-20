package model

type ReqMovie struct {
	Id          string
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Duration    string   `json:"duration" validate:"required"`
	Genres      []string `json:"genres" validate:"required"`
	Artists     []string `json:"artists" validate:"required"`
	WatchUrl    string
}
