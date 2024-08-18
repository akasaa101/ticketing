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
	TicketInsert(ticket models.Ticket) error
	TicketGetById(id int16) (models.Ticket, error)
}

func NewTicketService(repository repositories.TicketRepository) *TicketRepository {
	return &TicketRepository{repository}
}

func (tr TicketRepository) TicketInsert(ticket models.Ticket) error {
	err := tr.Repository.Insert(ticket) // This should be properly invoked
	if err != nil {
		return err
	}
	return nil
}

func (tr TicketRepository) TicketGetById(id int16) (models.Ticket, error) {
	result, err := tr.Repository.Get(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (tr TicketRepository) PurchaseTicket(id int16, quantity int) error {
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
