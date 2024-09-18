package entity

import(
	"time"
	"gorm.io/gorm"
)

type Payment struct{
	gorm.Model
	Date time.Time

	EmployeeID *uint
	Employee Employee `gorm:"foreignKey:EmployeeID"`
}