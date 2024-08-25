package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqliteDB(dbname string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Job{}, &Execution{})
	if err != nil {
		return nil, fmt.Errorf("could not migrate database: %w", err)
	}

	return db, nil
}
