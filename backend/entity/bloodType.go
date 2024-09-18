package entity

import (
	
	"gorm.io/gorm"
)

type BloodType struct{
	gorm.Model
	BloodGroup string

	Patients []Patient  `gorm:"foreignKey:BloodTypeID"`
	
}