package entity

import(
	"time"
	"gorm.io/gorm"
)

type Requisition struct{
	gorm.Model
	RequisitionQuantity int
	Time time.Time
	Note string

	EmployeeID *uint
	Employee Employee `gorm:"foreignKey:EmployeeID"`
}
	