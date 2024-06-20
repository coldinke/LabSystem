package models

import (
	"gorm.io/gorm"
)

const (
	Admin   = 1000
	Teacher = 1001
	Student = 1002
	Guest   = 1003
)

// Related table
type RoleRight struct {
	gorm.Model
	RoleID  uint `gorm:"index"`
	RightID uint `gorm:"index"`
}

type Role struct {
	gorm.Model
	RoleType int     `gorm:"column:role_type;type:int(11);not null" json:"roleType"`
	Rights   []Right `gorm:"many2many:roles_rights;" json:"rights"`
}

func (r *Role) TableName() string {
	return "roles"
}

func GetAllRoles(db *gorm.DB) ([]Role, error) {
	var roles []Role
	if err := db.Model(&Role{}).Preload("Rights").Select("id", "role_type").Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleByID(db *gorm.DB, id uint64) (Role, error) {
	var role Role
	result := db.Find(&role, id)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, nil
}

func UpdateRoleByID(db *gorm.DB, id uint64, data map[string]interface{}) error {
	role, err := GetRoleByID(db, id)
	if err != nil {
		return err
	}

	result := db.Model(&role).Updates(data)
	db.Model(&role).Association("Rights").Append(data["Rights"])
	return result.Error
}

func DeleteRoleByID(db *gorm.DB, id uint64) error {
	tx := db.Begin()
	var role Role
	if err := tx.First(&role, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&role).Association("Rights").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&role).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

type User struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(30);not null" json:"username"`
	Password  string `gorm:"column:password;type:varchar(40);not null" json:"-"`
	RoleID    uint   `gorm:"column:roleID;type:int(11);index;default:0" json:"role_id"`
	IsDefault bool   `gorm:"column:is_default;type:tinyint(1);default:0" json:"is_default"`
	Role      Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
}

func (u *User) TableName() string {
	return "users"
}

func GetUserList(db *gorm.DB) ([]User, error) {
	var users []User

	if err := db.Preload("Role").Preload("Role.Rights").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CheckUser(db *gorm.DB, username string, password string) (User, error) {
	var user User
	err := db.Preload("Role").Preload("Role.Rights").Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id uint64) (User, error) {
	var user User
	result := db.Find(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func UpdateUserByID(db *gorm.DB, id uint64, data map[string]interface{}) error {
	user, err := GetUserByID(db, id)
	if err != nil {
		return err
	}

	result := db.Model(&user).Updates(data)
	return result.Error
}

func CreateUser(db *gorm.DB, data map[string]interface{}) error {
	username := data["username"].(string)
	password := data["password"].(string)
	roleID := uint(data["roleID"].(int)) & 0xffffffff
	user := User{
		Username: username,
		Password: password,
		RoleID:   roleID,
	}
	result := db.Create(&user)
	return result.Error
}

func DeleteUserByID(db *gorm.DB, id uint64) error {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	result = db.Delete(&user)
	return result.Error
}
