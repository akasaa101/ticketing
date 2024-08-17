package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ticket struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Desc       string    `gorm:"type:text"`
	Allocation int       `gorm:"not null"`
	gorm.Model
}

type Tickets struct {
	Tickets []Ticket
}

func (ticket *Ticket) BeforeCreate(tx *gorm.DB) (err error) {
	ticket.ID = uuid.New()
	return
}
