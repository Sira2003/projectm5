package entity

import (
	"time"
	"gorm.io/gorm"
)
type DentalRecord struct{
	gorm.Model
	Date time.Time
	Description string
	Fees  float64
	Installment  float64
	NumberOfInstallments  int

	PatientID *uint
	Patient   Patient `gorm:"foreignKey:PatientID"`

	EmployeeID *uint
	Employee   Employee `gorm:"foreignKey:EmployeeID"`
}