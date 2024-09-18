package config

import (
	"fmt"
	"time"

	"example.com/projectm4/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Employee{},
		&entity.Patient{},
		&entity.DentalRecord{},
		&entity.Payment{},
		&entity.Requisition{},
		&entity.Schedule{},
		&entity.Gender{},
		&entity.JobPosition{},
		&entity.BloodType{},
	)

	GenderMale := entity.Gender{Sex: "ชาย"}
	GenderFemale := entity.Gender{Sex: "หญิง"}

	db.FirstOrCreate(&GenderMale, &entity.Gender{Sex: "ชาย"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Sex: "หญิง"})

	JobPositionDentist := entity.JobPosition{Job: "ทันตแพทย์"}
	JobPositionFinance := entity.JobPosition{Job: "เจ้าหน้าที่การเงิน"}
	JobPositionPatientService := entity.JobPosition{Job: "เจ้าหน้าที่บริการคนไข้"}

	db.FirstOrCreate(&JobPositionDentist, &entity.JobPosition{Job: "ทันตแพทย์"})
	db.FirstOrCreate(&JobPositionFinance, &entity.JobPosition{Job: "เจ้าหน้าที่การเงิน"})
	db.FirstOrCreate(&JobPositionPatientService, &entity.JobPosition{Job: "เจ้าหน้าที่บริการคนไข้"})

	BloodTypeA := entity.BloodType{BloodGroup: "A"}
	BloodTypeB := entity.BloodType{BloodGroup: "B"}
	BloodTypeAB := entity.BloodType{BloodGroup: "AB"}
	BloodTypeO := entity.BloodType{BloodGroup: "O"}

	db.FirstOrCreate(&BloodTypeA, &entity.BloodType{BloodGroup: "A"})
	db.FirstOrCreate(&BloodTypeB, &entity.BloodType{BloodGroup: "B"})
	db.FirstOrCreate(&BloodTypeAB, &entity.BloodType{BloodGroup: "AB"})
	db.FirstOrCreate(&BloodTypeO, &entity.BloodType{BloodGroup: "O"})
	//
	
	//

	hashedPassword, _ := HashPassword("123456")
	BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
	User := &entity.Employee{
		FirstName: "Software",
		LastName:  "Analysis",
		Email:     "sa@gmail.com",
		Password:  hashedPassword,
		Birthday:  BirthDay,
		GenderID:  1,
		JobPositionID: 3,
		Tel: "0000000000",
		Address: "บ้านเลขที่ 58 ถ.มิตรภาพ ต.ในเมือง อ.เมือง จ.นครราชสีมา 30000",
	}
	db.FirstOrCreate(User, &entity.Employee{
		Email: "sa@gmail.com",
	})

}

