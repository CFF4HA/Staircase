package database

import (
	"github.com/CFF4HA/Staircase/internal/engine/core"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	url string
	db  *gorm.DB
)

func Init(db_url string) {
	url = db_url

	conn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	core.Logger.Info("Database connection established")
	db = conn
}

func Database() *gorm.DB {
	if db == nil {

	}

	return db
}
