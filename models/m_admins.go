package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

// Admins :
type Admins struct {
	Model

	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhotoURL    string `json:"photo_url"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	UserRoles   string `json:"user_roles"`
	Verified    bool   `json:"verified"`
}

// ExistAdminByID :
func ExistAdminByID(id int) (bool, error) {
	var admin Admins
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if admin.ID > 0 {
		return true, nil
	}

	return false, fmt.Errorf("No admin with id %d", id)
}

// ExistAdminsByEmail : get member data by email
func ExistAdminsByEmail(email string) (Admins, error) {
	var member Admins
	err := db.Model(&Admins{}).Where("email = ? AND deleted_on = ?", email, 0).First(&member).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return Admins{}, err
	}
	if member.ID > 0 {
		return member, nil
	}
	return Admins{}, nil
}

// GetAdminTotal :
func GetAdminTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Admins{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetAdmin :
func GetAdmin(pageNum, pageSize int, maps interface{}) ([]Admins, error) {
	var (
		admin []Admins
		err   error
	)

	if pageSize > 0 && pageNum >= 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&admin).Error
	} else {
		err = db.Where(maps).Find(&admin).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return admin, nil
}

// AddAdmin :
func AddAdmin(data interface{}) error {

	var admin Admins
	err := mapstructure.Decode(data, &admin)
	if err != nil {
		return err
	}

	err = db.Create(&admin).Error
	if err != nil {
		return err
	}

	return nil
}

// EditAdmin :
func EditAdmin(id int, data interface{}) error {
	if err := db.Model(&Admins{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAdmin :
func DeleteAdmin(id int) error {
	if err := db.Where("id = ?", id).Delete(&Admins{}).Error; err != nil {
		return err
	}
	return nil
}
