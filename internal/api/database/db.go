package database

import (
	"github.com/CFF4HA/Staircase/internal/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDatabase(url string) {
	_db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = _db

	// This here will initialize the types
	// in the database via the AutoMigrate func.
	db.AutoMigrate(
		&types.DatabaseDatasource{},
		&types.DatabaseStaircase{},
		&types.DatabaseJob{},
		&types.DatabaseJobMetadata{},

		&types.User{},
	)
}

func Database() *gorm.DB {
	return db
}
