package entity
import (
"gorm.io/gorm"
)
type TStatus struct {
	gorm.Model
	TStatusName string
	Schedules [] Schedule `gorm:"foreignKey:TStatusID"`
}