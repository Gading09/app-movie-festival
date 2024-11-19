package container

import (
	"context"
	config "movie-festival/config"
	database "movie-festival/database"
	"time"

	featureMovie "movie-festival/domain/movie/feature"
	repoMovie "movie-festival/domain/movie/repository"

	featureUser "movie-festival/domain/user/feature"
	repoUser "movie-festival/domain/user/repository"

	"fmt"
	"log"

	"github.com/allegro/bigcache/v3"
	"gorm.io/gorm"
)

type Container struct {
	Database          *gorm.DB
	Cache             *bigcache.BigCache
	EnvironmentConfig config.EnvironmentConfig
	MovieFeature      featureMovie.MovieFeature
	UserFeature       featureUser.UserFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading config...")
	config, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Loading database...")
	db, err := database.DatabaseConnection(config.Database)
	if err != nil {
		log.Panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	cache, err := bigcache.New(ctx, bigcache.DefaultConfig(1*time.Hour))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading repository's...")
	repoMovie := repoMovie.NewMovieRepository(db)
	repoUser := repoUser.NewUserRepository(db)

	fmt.Println("Loading feature's...")
	featureMovie := featureMovie.NewMovieFeature(repoMovie)
	featureUser := featureUser.NewUserFeature(repoUser, cache)

	return Container{
		Database:          db,
		Cache:             cache,
		EnvironmentConfig: config,
		MovieFeature:      featureMovie,
		UserFeature:       featureUser,
	}
}
