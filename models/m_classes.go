package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

// Classes :
type Classes struct {
	Base `mapstructure:",squash"`

	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Headline    string `json:"headline"`
}

// ExistClassByID :
func ExistClassByID(id int64) (bool, error) {
	var class Classes
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&class).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if class.ID > 0 {
		return true, nil
	}

	return false, fmt.Errorf("No class with id %d", id)
}

// GetClassTotal :
func GetClassTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Classes{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetClasses :
func GetClasses(pageNum, pageSize int, maps interface{}) ([]Classes, error) {
	var (
		class []Classes
		err   error
	)

	dbClass := db.Where(maps)

	if pageSize > 0 && pageNum >= 0 {
		dbClass = dbClass.Offset(pageNum).Limit(pageSize)
	}

	err = dbClass.Order("id desc").Find(&class).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return class, nil
}

// AddClass :
func AddClass(data interface{}) error {
	fmt.Println(data)
	var class Classes
	err := mapstructure.Decode(data, &class)
	if err != nil {
		return err
	}

	fmt.Println(class)
	err = db.Create(&class).Error
	if err != nil {
		return err
	}

	return nil
}

// EditClass :
func EditClass(id int64, data interface{}) error {
	if err := db.Model(&Classes{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteClass :
func DeleteClass(id int64, uuid string) error {
	if err := db.Model(&Classes{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(map[string]interface{}{"deleted_by": uuid}).Error; err != nil {
		return err
	}
	if err := db.Where("id = ?", id).Delete(&Classes{}).Error; err != nil {
		return err
	}
	return nil
}
