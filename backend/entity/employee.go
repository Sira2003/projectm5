package entity

import (
	"time"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
	Address string
	Tel string `gorm:"unique"`
	Password string
	Email string `gorm:"unique"`
	
	Payments []Payment `gorm:"foreignKey:EmployeeID"`
	DentalRecords []DentalRecord `gorm:"foreignKey:EmployeeID"`
	Requisitions []Requisition `gorm:"foreignKey:EmployeeID"`
	Restocks []Restock `gorm:"foreignKey:EmployeeID"`
	
	GenderID uint
	Gender   Gender `gorm:"foreignKey:GenderID"`

	JobPositionID uint
	JobPosition   JobPosition `gorm:"foreignKey:JobPositionID"`
}

	// 1 member เป็นเจ้าของได้หลาย Playlist
	//Playlists []Playlist `gorm:"foreignKey:MemberID"`