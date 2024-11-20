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

type ReqUpdateMovie struct {
	Id          string
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Duration    string   `json:"duration" validate:"required"`
	Genres      []string `json:"genres" validate:"required"`
	Artists     []string `json:"artists" validate:"required"`
	WatchUrl    string   `json:"url" validate:"required"`
}

type ReqGetListMovie struct {
	Page  int `json:"page" validate:"required"`
	Limit int `json:"limit" validate:"required"`
}
