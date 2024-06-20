package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"log"
	"testing"
)

// Role 和 Right 结构体定义保持不变

func TestRole(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Role{})

	rolesRightsMap := []struct {
		RoleType int
		RightIDs []uint
	}{
		{
			RoleType: models.Admin,
			RightIDs: []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			RoleType: models.Teacher,
			RightIDs: []uint{1, 2, 3, 8, 9, 10, 13},
		},
		{
			RoleType: models.Student,
			RightIDs: []uint{1, 10, 13},
		},
		{
			RoleType: models.Guest,
			RightIDs: []uint{1},
		},
	}

	for _, item := range rolesRightsMap {
		role := models.Role{
			RoleType: item.RoleType,
		}
		db.Create(&role)
		var rights []models.Right
		db.Find(&rights, item.RightIDs)
		db.Model(&role).Association("Rights").Replace(rights) // 替换角色的权限为当前查询到的权限

	}

	target := int64(len(rolesRightsMap))
	count := int64(0)
	result := db.Model(&models.Role{}).Count(&count)
	if result.Error != nil {
		log.Fatal("Query Error", result.Error)
	}
	assert.Equal(t, target, count)
}
