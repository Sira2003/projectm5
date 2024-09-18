package entity

import (
	
	"gorm.io/gorm"
)

type Gender struct{
	gorm.Model
	Sex string

	Employees []Employee  `gorm:"foreignKey:GenderID"`
	Patients []Patient  `gorm:"foreignKey:GenderID"`
	
}