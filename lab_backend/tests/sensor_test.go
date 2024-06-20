package tests

//
//func TestSensor(t *testing.T) {
//	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	db.AutoMigrate(&models.Sensor{})
//
//	sensor := models.Sensor{
//		Name:  "Temperature",
//		Type:  3,
//		LabID: 1,
//	}
//
//	log.Printf("%d", sensor.ID)
//
//	db.Create(&sensor)
//
//	db.AutoMigrate(&models.SensorData{})
//
//	data := models.SensorData{
//		Value:     20.5,
//		Timestamp: time.Now(),
//		SensorID:  sensor.ID,
//	}
//
//	db.Create(&data)
//}
