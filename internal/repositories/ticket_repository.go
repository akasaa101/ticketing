package repositories

import (
	"errors"
	"fmt"
	"github.com/akasaa101/ticketing/internal/database"
	"github.com/akasaa101/ticketing/internal/models"
	"gorm.io/gorm"
	"log"
)

type DbInstance struct {
	DB *gorm.DB
}

type TicketRepository interface {
	Insert(ticket models.Ticket) (models.Ticket, error)
	Get(id int16) (models.Ticket, error)
	Update(ticket models.Ticket) error
}

func (tdb DbInstance) Insert(ticket models.Ticket) (models.Ticket, error) {
	if result := tdb.DB.Create(&ticket); result.Error != nil {
		log.Printf("ticketRepository Insert error : %s", result.Error)
		return models.Ticket{}, result.Error
	}
	return ticket, nil
}

func (bdb DbInstance) Get(id int16) (models.Ticket, error) {
	var ticket models.Ticket
	if result := bdb.DB.First(&ticket, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("ticketRepository Get error: record not found for id %d", id)
			return models.Ticket{}, fmt.Errorf("ticket with ID %d not found", id)
		}
		log.Printf("ticketRepository Get error: %s", result.Error)
		return models.Ticket{}, result.Error
	}
	return ticket, nil
}

func (tdb DbInstance) Update(ticket models.Ticket) error {
	if result := tdb.DB.Save(&ticket); result.Error != nil {
		log.Printf("ticketRepository Update error : %s", result.Error)
		return result.Error
	}
	return nil
}

func NewTicketRepositoryDB() *DbInstance {
	return &DbInstance{DB: database.DB.Db}
}
