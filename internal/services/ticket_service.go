package services

import (
	"errors"
	"github.com/akasaa101/ticketing/internal/models"
	"github.com/akasaa101/ticketing/internal/repositories"
)

type TicketRepository struct {
	Repository repositories.TicketRepository
}

type TicketService interface {
	TicketInsert(ticket models.Ticket) (models.Ticket, error)
	TicketGetById(id int16) (models.Ticket, error)
	PurchaseTicket(id int16, quantity int) error
}

func NewTicketService(repository repositories.TicketRepository) *TicketRepository {
	return &TicketRepository{repository}
}

func (tr TicketRepository) TicketInsert(ticket models.Ticket) (models.Ticket, error) {
	insertedTicket, err := tr.Repository.Insert(ticket)
	if err != nil {
		return models.Ticket{}, err
	}
	return insertedTicket, nil
}

func (tr TicketRepository) TicketGetById(id int16) (models.Ticket, error) {
	ticket, err := tr.Repository.Get(id)
	if err != nil {
		return models.Ticket{}, err
	}
	return ticket, nil
}
func (tr TicketRepository) PurchaseTicket(id int16, quantity int) error {
	if quantity < 0 {
		return errors.New("non-positive quantity")
	}
	ticket, err := tr.Repository.Get(id)
	if err != nil {
		return err
	}

	if ticket.Allocation < quantity {
		return errors.New("insufficient allocation")
	}

	ticket.Allocation -= quantity
	return tr.Repository.Update(ticket)
}
