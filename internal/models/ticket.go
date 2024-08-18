package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Id         int    `json:"id" gorm:"primary_key;"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}
