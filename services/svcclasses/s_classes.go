package svcclasses

import (
	"kusnandartoni/starter/models"

	"github.com/fatih/structs"
)

// Classes :
type Classes struct {
	ID         int64  `json:"id,omitempty" structs:"id,omitempty"`
	CreatedBy  string `json:"created_by,omitempty" structs:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty" structs:"modified_by,omitempty"`
	DeletedBy  string `json:"deleted_by,omitempty" structs:"deleted_by,omitempty"`

	ImageURL    string `json:"image_url,omitempty" structs:"image_url,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	Headline    string `json:"headline,omitempty" structs:"headline,omitempty"`

	PageNum  int `json:"page_num,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// ExistByID :
func (s *Classes) ExistByID() (bool, error) {
	return models.ExistClassByID(s.ID)
}

// Count :
func (s *Classes) Count() (int, error) {
	return models.GetClassTotal(s.getMaps())
}

// Add :
func (s *Classes) Add() error {
	return models.AddClass(s)
}

// GetAll :
func (s *Classes) GetAll() ([]models.Classes, error) {
	class, err := models.GetClasses(s.PageNum, s.PageSize, s.getMaps())
	if err != nil {
		return nil, err
	}

	return class, nil
}

// Edit :
func (s *Classes) Edit() error {
	data := structs.Map(s)
	return models.EditClass(s.ID, data)
}

// Delete :
func (s *Classes) Delete() error {
	return models.DeleteClass(s.ID, s.DeletedBy)
}

func (s *Classes) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if s.ID != -1 {
		maps["id"] = s.ID
	}
	return maps
}
