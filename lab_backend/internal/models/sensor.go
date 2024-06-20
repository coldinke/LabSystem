package models

import (
	"gorm.io/gorm"
)

type SensorData struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(50); not null"`
	Fire        int    `gorm:"column:fire;type:tinyint(3);not null"`
	Smoke       int    `gorm:"column:smoke;type:tinyint(3);not null"`
	Temperature int    `gorm:"column:temperature;type:int(11);not null"`
	Humidity    int    `gorm:"column:humidity;type:int(11);not null"`
	LabID       uint
	Lab         Lab `json:"-"`
}

func (d *SensorData) TableName() string {
	return "sensor_data"
}

func SaveSensorData(db *gorm.DB, data SensorData) error {
	if err := db.First(&data.Lab, data.LabID).Error; err != nil {
		return err
	}
	if result := db.Create(&data); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLatestSensorDataByLabID(db *gorm.DB, labID uint) (SensorData, error) {
	var sensorData SensorData
	result := db.Where("lab_id = ?", labID).Order("created_at desc").First(&sensorData)
	if result.Error != nil {
		return SensorData{}, result.Error
	}
	return sensorData, nil
}

func GetSensorList(db *gorm.DB, labID uint64) ([]SensorData, error) {
	var datas []SensorData
	err := db.Model(SensorData{}).Limit(10).Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	return datas, nil
}
