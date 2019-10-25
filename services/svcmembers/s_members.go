package svcmembers

import (
	"errors"
	"kusnandartoni/starter/models"
	"kusnandartoni/starter/pkg/util"

	"github.com/fatih/structs"
)

// Members :
type Members struct {
	ID       int
	Email    string `json:"email" structs:"email,omitempty"`
	Password string `json:"password" structs:"password,omitempty"`
	FullName string `json:"full_name" structs:"full_name,omitempty"`
	PhotoURL string `json:"photo_url" structs:"photo_url,omitempty"`
	Verified bool   `json:"verified" structs:"verified,omitempty"`

	PageNum  int `json:"page_num,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// Add : create new member
func (m *Members) Add() error {
	return models.AddMembers(m)
}

// GetByEmail :
func (m *Members) GetByEmail() (models.Members, error) {
	return models.ExistMembersByEmail(m.Email)
}

// GetAll :
func (m *Members) GetAll() ([]models.Members, error) {
	var class []models.Members

	class, err := models.GetMembers(m.PageNum, m.PageSize, m.getMaps())
	if err != nil {
		return nil, err
	}

	return class, nil
}

// Edit :
func (m *Members) Edit() error {
	data := structs.Map(m)
	return models.EditMembers(m.ID, data)
}

// Identify : get member auth by email and password
func (m *Members) Identify() (models.Members, error) {
	member, err := models.ExistMembersByEmail(m.Email)
	if err != nil {
		return models.Members{}, errors.New("Invalid Email")
	}

	_, err = util.Compare(member.Password, m.Password)
	if err != nil {
		return models.Members{}, errors.New("Invalid Password")
	}
	member.Password = ""
	return member, nil
}

func (m *Members) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if m.ID != -1 {
		maps["id"] = m.ID
	}
	return maps
}
