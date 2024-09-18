package entity

import (
	"time"
	"gorm.io/gorm"
)

type Schedule struct{
	gorm.Model
	Date time.Time

	// TreatmentID uint
	// TStatusID uint

	PatientID *uint
	Patient   Patient `gorm:"foreignKey:PatientID"`
}