package database

import (
	"time"
)

type Job struct {
	Title     string    `gorm:"not null,uniqueIndex" json:"title"`
	Image     string    `gorm:"default:bash" json:"image"`
	Script   string     `gorm:"not null" json:"script"`
	Variables string    `gorm:"type:json"`
	LastSync  time.Time `gorm:"not null"`
}

type Execution struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"not null"`
	Image         string    `gorm:"default:bash"`
	Script       string    `gorm:"not null"`
	Variables     string    `gorm:"type:json"`
	Hash          string    `gorm:"not null"`
	RunAt         time.Time `gorm:"not null"`
	Status        int       `gorm:"not null"` // You can replace int with a custom enum type if needed
	Output        string
	Artifacts     string `gorm:"type:json"`
	StorageDriver string `gorm:"not null"`
}
