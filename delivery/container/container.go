package container

import (
	config "movie-festival/config"
	database "movie-festival/database"

	"fmt"
	"log"

	"gorm.io/gorm"
)

type Container struct {
	Database          *gorm.DB
	EnvironmentConfig config.EnvironmentConfig
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

	return Container{
		Database:          db,
		EnvironmentConfig: config,
	}
}
