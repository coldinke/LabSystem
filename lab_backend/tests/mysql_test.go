package tests

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"testing"
)

func TestMysql(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&product, 1)

	// Test Sensor

	db.AutoMigrate(&models.SensorData{
		Name:        "test",
		LabID:       1,
		Fire:        1,
		Smoke:       1,
		Flying:      1,
		Temperature: 1,
		Humidity:    1,
	})

	//var lab models.Lab
	////var data models.SensorData
	//
	//db.First(&lab, 1)
	//
	////
	//db.Where("lab_id = ?", lab.ID).Find(&sensor)
	//
	//for i := 0; i < 3; i++ {
	//	log.Printf("%d", sensor[i].ID)
	//	db.Preload("Data").First(&sensor[i], "id = ?", sensor[i].ID)
	//	log.Printf("%+v", sensor[i].Data)
	//}

	////db.Model(&sensor).Where("sensor_id = ?", sensor.ID).Association("SensorData").Find(&data)
	//
	//db.Preload("Data").First(&sensor, "id = ?", sensor.ID)
	//log.Printf("%+v", sensor.Data.Value)

	//db.First(&data, sensor.ID)

	//log.Printf("Lab %+v", data)
}
