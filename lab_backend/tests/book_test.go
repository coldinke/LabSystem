package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"log"
	"testing"
)

func TestBook(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(models.Book{})

	book := models.Book{
		LabID:    1,
		Reason:   "test reason",
		Username: "Vinci",
	}

	db.Create(&book)

	count := int64(0)
	result := db.Model(&models.Lab{}).Count(&count)
	if result.Error != nil {
		log.Fatal("Query Error", result.Error)
	}
	assert.Equal(t, result, count)
}
