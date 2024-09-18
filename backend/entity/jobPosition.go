package entity

import (
	
	"gorm.io/gorm"
)

type JobPosition struct{
	gorm.Model
	Job string

	Employees [] Employee  `gorm:"foreignKey:JobPositionID"`
	
	
}