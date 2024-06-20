package models

import (
	"gorm.io/gorm"
	"time"
)

// LabType
const (
	ChemistryLabType = 1 + iota
	BiologicalLabType
	RadiationLabType
	ElectromechanicalLabType
	OtherLabType
)

// RiskType
const (
	Level1 = 1 + iota
	Level2
	Level3
	Level4
)

// ColegeType
const (
	Info = 1 + iota
	Math
)

type Lab struct {
	gorm.Model
	LabName     string `gorm:"column:lab_name; type:varchar(50);" json:"lab_name"`
	LabType     int    `gorm:"column:lab_type; type:int(11);" json:"lab_type"`
	RiskType    int    `gorm:"column:risk_type; type:int(11);" json:"risk_type"`
	CollegeType int    `gorm:"column:college_type; type:int(11);" json:"college_type"`
	Capacity    int    `gorm:"column:capacity; type:int(11)" json:"capacity"`
}

func (l *Lab) TableName() string {
	return "labs"
}

func GetLabList(db *gorm.DB) ([]Lab, error) {
	var labs []Lab
	err := db.Model(Lab{}).Find(&labs).Error
	if err != nil {
		return nil, err
	}
	return labs, nil
}

func GetLabByID(db *gorm.DB, id uint64) (Lab, error) {
	var lab Lab
	result := db.Find(&lab, id)
	if result.Error != nil {
		return Lab{}, result.Error
	}
	return lab, nil
}

func UpdateLabByID(db *gorm.DB, id uint64, data map[string]interface{}) error {
	lab, err := GetLabByID(db, id)
	if err != nil {
		return err
	}

	result := db.Model(&lab).Updates(data)
	return result.Error
}

func CreateLab(db *gorm.DB, data map[string]interface{}) error {
	labname := data["lab_name"].(string)
	labtype := data["lab_type"].(int)
	risktype := data["risk_type"].(int)
	collegetype := data["college_type"].(int)
	capacity := data["capacity"].(int)

	lab := Lab{
		LabName:     labname,
		LabType:     labtype,
		RiskType:    risktype,
		CollegeType: collegetype,
		Capacity:    capacity,
	}

	result := db.Create(&lab)
	return result.Error

}

func DeleteLabByID(db *gorm.DB, id uint64) error {
	var lab Lab
	result := db.First(&lab, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&lab)
	return result.Error
}

type LabWithSensorData struct {
	Lab         Lab          `json:"lab"`
	SensorDatas []SensorData `json:"sensor_datas"`
}

func GetLabDetailByID(db *gorm.DB, id uint64) (LabWithSensorData, error) {
	lab, err := GetLabByID(db, id)
	if err != nil {
		return LabWithSensorData{}, err
	}

	var sensorDatas []SensorData
	sensorData, _ := GetLatestSensorDataByLabID(db, lab.ID)
	sensorDatas = append(sensorDatas, sensorData)
	return LabWithSensorData{
		Lab:         lab,
		SensorDatas: sensorDatas,
	}, nil
}

const (
	Accept     = 0
	Reject     = 1
	Processing = 2
)

type Book struct {
	gorm.Model
	LabID    uint      `gorm:"column:lab_id;default:0"`
	ClassID  int       `gorm:"column:class_id;type:int(11)"`
	Reason   string    `gorm:"column:book_reason;type:varchar(100);not null"`
	Username string    `gorm:"column:book_username;type:varchar(20);not null"`
	State    int       `gorm:"column:book_state;type:int;default:2"`
	BookTime time.Time `gorm:"column:book_time;type:datetime"`
}

func (b *Book) TableName() string {
	return "books"
}

func GetBookList(db *gorm.DB) ([]Book, error) {
	var books []Book
	err := db.Model(Book{}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(db *gorm.DB, book Book) error {
	result := db.Create(&book)
	return result.Error
}

func GetBookListByUsername(db *gorm.DB, username string) ([]Book, error) {
	var books []Book
	err := db.Model(Book{}).Where("book_username = ? ", username).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookListByLabID(db *gorm.DB, labID uint64, bookTime time.Time) ([]Book, error) {
	var books []Book
	err := db.Model(Book{}).Where("lab_id = ? AND book_time = ? AND book_state = ?", labID, bookTime, 0).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(db *gorm.DB, id uint64) (Book, error) {
	var book Book
	result := db.Find(&book, id)
	if result.Error != nil {
		return Book{}, result.Error
	}
	return book, nil
}

func DeleteBookByID(db *gorm.DB, id uint64) error {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&book)
	return result.Error
}

func GetBookListByStatus(db *gorm.DB, status int) ([]Book, error) {
	var books []Book
	err := db.Model(Book{}).Where("book_state = ? AND book_time > ?", status, time.Now()).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func UpdateBookByID(db *gorm.DB, id uint64, state int) error {
	book, err := GetBookByID(db, id)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"book_state": state,
	}

	result := db.Model(&book).Updates(data)
	return result.Error
}
