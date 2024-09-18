package entity

import (
	"time"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
	Weight	int
	Height int	
	// BloodType  string
	DrugAllergy string
	Chronicdisease  string
	Tel string `gorm:"unique"`
	
	Schedules [] Schedule  `gorm:"foreignKey:PatientID"`
	DentalRecords [] DentalRecord  `gorm:"foreignKey:PatientID"`

	GenderID uint
	Gender   Gender `gorm:"foreignKey:GenderID"`

	BloodTypeID uint
	BloodType   BloodType `gorm:"foreignKey:BloodTypeID"`
	
}

	// 1 member เป็นเจ้าของได้หลาย Playlist
	//Playlists []Playlist `gorm:"foreignKey:MemberID"`