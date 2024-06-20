package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	testUser1 := models.User{
		Username:  "Vinci",
		Password:  "hello1",
		RoleID:    1,
		IsDefault: true,
	}

	testUser2 := models.User{
		Username:  "teacher",
		Password:  "hello1",
		RoleID:    2,
		IsDefault: false,
	}

	db.Create(&testUser1)
	db.Create(&testUser2)

	count := int64(0)
	result := db.Model(&models.User{}).Count(&count)
	if result.Error != nil {
		log.Fatal("Query Error", result.Error)
	}
	assert.Equal(t, result, count)
}
