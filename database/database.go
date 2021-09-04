package database

import (
	"github.com/GolangSriLanka/go-puso/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetEnv("database.URL")), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
	}
	log.Info("Database connected")
	return db
}
