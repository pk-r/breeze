package database

import "gorm.io/gorm"

type JobRepository interface {
	Sync([]Job) error
}

type GormJobRepository struct {
	DB *gorm.DB
}

func (gjr GormJobRepository) Sync(jobs []Job) error {
	tx := gjr.DB.Begin()

	if err := tx.Exec("DELETE FROM jobs").Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return err
	}

	// Insert new jobs
	if err := tx.Create(&jobs).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
