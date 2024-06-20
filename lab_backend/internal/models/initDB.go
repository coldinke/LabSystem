package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = initDB()

func initDB() *gorm.DB {
	dsn := "root:qweasdzxc@tcp(mysql_labSystem.orb.local:3306)/lab_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect the Mysql: %s", err)
	}

	//AutoMigrate(db)
	//initData(db)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Right{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Lab{})
	db.AutoMigrate(&SensorData{})
	db.AutoMigrate(&Book{})
}

func initData(db *gorm.DB) {

	// Rights
	initialRights := []Right{
		{Title: "首页", Path: "/index", Icon: "HomeFilled"},
		{Title: "用户管理", Path: "/user-manage", Icon: "User"},
		{Title: "用户列表", Path: "/user-manage/userList", Icon: "List", ParentId: 2},
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

	db.Create(&initialRights)

	// Roles
	rolesRightsMap := []struct {
		RoleType int
		RightIDs []uint
	}{
		{
			RoleType: Admin,
			RightIDs: []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			RoleType: Teacher,
			RightIDs: []uint{1, 2, 3, 8, 9, 11, 12},
		},
		{
			RoleType: Student,
			RightIDs: []uint{1, 9, 10},
		},
		{
			RoleType: Guest,
			RightIDs: []uint{1},
		},
	}

	for _, item := range rolesRightsMap {
		role := Role{
			RoleType: item.RoleType,
		}
		db.Create(&role)
		var rights []Right
		db.Find(&rights, item.RightIDs)
		db.Model(&role).Association("Rights").Replace(rights)
	}

	// Users
	initialUsers := []User{
		{
			Username:  "Vinci",
			Password:  "hello1",
			RoleID:    1,
			IsDefault: true,
		},
		{
			Username:  "teacher",
			Password:  "hello1",
			RoleID:    2,
			IsDefault: false,
		},
	}

	db.Create(initialUsers)

	// Labs
	initialLabs := []Lab{
		{
			LabName:     "testLab",
			LabType:     BiologicalLabType,
			RiskType:    Level2,
			CollegeType: Math,
			Capacity:    100,
		},
		{
			LabName:     "testLab2",
			LabType:     ElectromechanicalLabType,
			RiskType:    Level1,
			CollegeType: ElectromechanicalLabType,
			Capacity:    50,
		},
	}

	db.Create(initialLabs)
}
