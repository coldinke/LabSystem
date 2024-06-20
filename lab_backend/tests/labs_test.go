package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"log"
	"testing"
)

func TestLab(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Lab{})

	testLab1 := models.Lab{
		LabName:     "testLab",
		LabType:     models.BiologicalLabType,
		RiskType:    models.Level2,
		CollegeType: models.Math,
		Capacity:    100,
	}

	db.Create(&testLab1)

	count := int64(0)
	result := db.Model(&models.Lab{}).Count(&count)
	if result.Error != nil {
		log.Fatal("Query Error", result.Error)
	}
	assert.Equal(t, 1, count)
}
