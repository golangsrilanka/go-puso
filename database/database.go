package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/GolangSriLanka/go-puso/config"
)

var db *gorm.DB
var err error

func Database() *gorm.DB {
	if db == nil {
		db, err = gorm.Open(postgres.Open(config.GetEnv("database.URL")), &gorm.Config{})

		if err != nil {
			log.Panic("failed to connect database")
		}

		log.Info("database connected")

		return db
	}

	return db
}
