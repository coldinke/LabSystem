package tests

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab_backend/internal/models"
	"log"
	"testing"
)

func TestRighst(t *testing.T) {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Right{})

	initialRights := []models.Right{
		{Title: "首页", Path: "/index", Icon: "HomeFilled"},
		{Title: "用户管理", Path: "/user-manage", Icon: "User"},
		{Title: "用户列表", Path: "/user-manage/list", Icon: "List", ParentId: 2},
		{Title: "权限管理", Path: "/right-manage", Icon: "Key"},
		{Title: "角色列表", Path: "/right-manage/roleList", Icon: "List", ParentId: 4},
		{Title: "权限列表", Path: "/right-manage/rightList", Icon: "List", ParentId: 4},
		{Title: "实验室管理", Path: "/lab-manage", Icon: "OfficeBuilding"},
		{Title: "实验室列表", Path: "/lab-manage/labList", Icon: "List", ParentId: 7},
		{Title: "预约管理", Path: "/book-manage", Icon: "UploadFilled"},
		{Title: "审核管理", Path: "/book-manage/auditList", Icon: "List", ParentId: 9},
		{Title: "预约列表", Path: "/book-manage/bookList", Icon: "List", ParentId: 9},
		{Title: "预约实验室", Path: "/book-manage/bookLab", Icon: "Key", ParentId: 9},
	}

	for _, right := range initialRights {
		db.Create(&right) // 注意这里传递的是 right 的指针
	}

	target := int64(len(initialRights))
	count := int64(0)
	result := db.Model(&models.Right{}).Count(&count)
	if result.Error != nil {
		log.Fatal("Query Error", result.Error)
	}

	assert.Equal(t, target, count)
}
