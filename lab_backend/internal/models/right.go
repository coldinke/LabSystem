package models

import (
	"gorm.io/gorm"
	"log"
)

type Right struct {
	gorm.Model
	Title    string `gorm:"column:title;type:varchar(20);not null;" json:"title"`
	Path     string `gorm:"column:path;type:varchar(100);not null;" json:"path"`
	Icon     string `gorm:"column:icon;type:varchar(100);not null" json:"icon"`
	ParentId int    `gorm:"column:parent_id;type:int;" json:"parent_id"`
}

func (r *Right) TableName() string {
	return "rights"
}

func getAllRight(db *gorm.DB) ([]*Right, error) {
	var rights []*Right
	if err := db.Find(&rights).Error; err != nil {
		return nil, err
	}
	return rights, nil
}

func GetRightByID(db *gorm.DB, id uint64) (Right, error) {
	var right Right
	result := db.Find(&right, id)
	if result.Error != nil {
		return Right{}, result.Error
	}
	return right, result.Error
}

func UpdateRightByID(db *gorm.DB, id uint64, data map[string]interface{}) error {
	right, err := GetRightByID(db, id)
	if err != nil {
		return err
	}

	result := db.Model(&right).Updates(data)
	return result.Error
}

func DeleteRightByID(db *gorm.DB, id uint64) error {
	tx := db.Begin()
	var right Right
	print(1)
	if err := tx.First(&right, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	print(2)
	if err := tx.Exec("DELETE FROM roles_rights WHERE right_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	print(3)
	if err := tx.Delete(&right).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

type MenuItem struct {
	Id       uint       `json:"id"`
	Title    string     `json:"title"`
	Path     string     `json:"path"`
	Icon     string     `json:"icon"`
	Children []MenuItem `json:"children"`
}

func GetMenuList() ([]MenuItem, error) {
	var menuList []MenuItem

	rights, err := getAllRight(DB)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rights {
		if r.ParentId == 0 {
			item := MenuItem{
				Id:    r.ID,
				Title: r.Title,
				Path:  r.Path,
				Icon:  r.Icon,
			}
			children := getChildren(rights, r.ID)
			if children == nil {
				children = []MenuItem{}
			}
			item.Children = children
			menuList = append(menuList, item)
		}
	}

	return menuList, nil
}

func getChildren(rights []*Right, parentId uint) []MenuItem {
	var children []MenuItem

	for _, r := range rights {
		if r.ParentId == int(parentId) {
			item := MenuItem{
				Id:    r.ID,
				Title: r.Title,
				Path:  r.Path,
				Icon:  r.Icon,
			}
			childrenItems := getChildren(rights, r.ID)
			if childrenItems == nil {
				childrenItems = []MenuItem{}
			}
			item.Children = childrenItems

			children = append(children, item)
		}
	}
	return children
}
