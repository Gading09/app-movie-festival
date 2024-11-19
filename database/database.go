package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Name     string
	Username string
	Password string
	Port     string
}

func DatabaseConnection(config DatabaseConfig) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	return gorm.Open(mysql.Open(dataSource), &gorm.Config{})
}
