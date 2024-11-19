package feature

import (
	"movie-festival/domain/user/repository"

	"github.com/allegro/bigcache/v3"
)

type UserFeature interface {
}

type userFeature struct {
	Repository repository.UserRepository
	Cache      *bigcache.BigCache
}

func NewUserFeature(repository repository.UserRepository, cache *bigcache.BigCache) UserFeature {
	return &userFeature{
		Repository: repository,
		Cache:      cache,
	}
}
