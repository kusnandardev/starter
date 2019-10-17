package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

// Mentors : models struct
type Mentors struct {
	Model

	Email       string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhotoURL    string `json:"photo_url"`
	Expertise   string `json:"expertise"`
	Biography   string `json:"biography"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Birthday    string `json:"birthday"`
	Address     string `json:"address"`
	Verified    bool   `json:"verified"`
	Website     string `json:"website"`
	Institution string `json:"institution"`
}

// ExistMentorByID :
func ExistMentorByID(id int) (bool, error) {
	var mentor Mentors
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&mentor).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if mentor.ID > 0 {
		return true, nil
	}

	return false, fmt.Errorf("No mentor with id %d", id)
}

// ExistMentorsByEmail : get member data by email
func ExistMentorsByEmail(email string) (Mentors, error) {
	var member Mentors
	err := db.Model(&Mentors{}).Where("email = ? AND deleted_on = ?", email, 0).First(&member).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return Mentors{}, err
	}
	if member.ID > 0 {
		return member, nil
	}
	return Mentors{}, nil
}

// GetMentorTotal :
func GetMentorTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Mentors{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetMentor :
func GetMentor(pageNum, pageSize int, maps interface{}) ([]Mentors, error) {
	var (
		mentor []Mentors
		err    error
	)

	if pageSize > 0 && pageNum >= 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&mentor).Error
	} else {
		err = db.Where(maps).Find(&mentor).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return mentor, nil
}

// AddMentor :
func AddMentor(data interface{}) error {

	var mentor Mentors
	err := mapstructure.Decode(data, &mentor)
	if err != nil {
		return err
	}

	err = db.Create(&mentor).Error
	if err != nil {
		return err
	}

	return nil
}

// EditMentor :
func EditMentor(id int, data interface{}) error {
	if err := db.Model(&Mentors{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// DeleteMentor :
func DeleteMentor(id int) error {
	if err := db.Where("id = ?", id).Delete(&Mentors{}).Error; err != nil {
		return err
	}
	return nil
}
