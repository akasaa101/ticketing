package models

import (
	"errors"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Id         int    `json:"id" gorm:"primary_key;"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Allocation < 0 {
		return errors.New("allocation cannot be negative")
	}
	return nil
}

func (t *Ticket) BeforeUpdate(tx *gorm.DB) (err error) {
	if t.Allocation < 0 {
		return errors.New("allocation cannot be negative")
	}
	return nil
}
