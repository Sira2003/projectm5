package entity

import(
	"time"
	"gorm.io/gorm"
)

type Restock struct{
	gorm.Model
	RestockQuantity int
	RestockDate time.Time
	Note string

	EmployeeID *uint
	Employee Employee `gorm:"foreignKey:EmployeeID"`
}