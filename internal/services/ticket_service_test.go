package services_test

import (
	"errors"
	"testing"

	"github.com/akasaa101/ticketing/internal/models"

	"github.com/akasaa101/ticketing/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(ticket models.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

func (m *MockRepository) Get(id int16) (models.Ticket, error) {
	args := m.Called(id)
	return args.Get(0).(models.Ticket), args.Error(1)
}

func (m *MockRepository) Update(ticket models.Ticket) error {
	args := m.Called(ticket)
	return args.Error(0)
}

func TestTicketService_TicketInsert(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.NewTicketService(mockRepo)

	ticket := models.Ticket{
		Name:       "example",
		Desc:       "sample description",
		Allocation: 100,
	}

	t.Run("Success", func(t *testing.T) {
		mockRepo.On("Insert", ticket).Return(nil)

		err := service.TicketInsert(ticket)

		assert.NoError(t, err)

		mockRepo.AssertCalled(t, "Insert", ticket)
		mockRepo.AssertExpectations(t)
	})

}

func TestTicketService_GetTicketById(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.NewTicketService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		ticketID := int16(1)
		expectedTicket := models.Ticket{
			Id:         int(ticketID),
			Name:       "example",
			Desc:       "sample description",
			Allocation: 100,
		}

		mockRepo.On("Get", ticketID).Return(expectedTicket, nil)

		ticket, err := service.TicketGetById(ticketID)

		assert.NoError(t, err)
		assert.Equal(t, expectedTicket, ticket)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Ticket Not Found", func(t *testing.T) {
		ticketID := int16(2)

		mockRepo.On("Get", ticketID).Return(models.Ticket{}, errors.New("ticket not found"))

		ticket, err := service.TicketGetById(ticketID)

		assert.Error(t, err)
		assert.Equal(t, "ticket not found", err.Error())
		assert.Equal(t, models.Ticket{}, ticket)
		mockRepo.AssertExpectations(t)
	})
}

func TestTicketService_PurchaseTicket(t *testing.T) {
	mockRepo := new(MockRepository)
	service := services.NewTicketService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		ticketID := int16(1)
		initialTicket := models.Ticket{
			Id:         int(ticketID),
			Name:       "example",
			Desc:       "sample description",
			Allocation: 100,
		}
		updatedTicket := initialTicket
		updatedTicket.Allocation = 90

		mockRepo.On("Get", ticketID).Return(initialTicket, nil)
		mockRepo.On("Update", updatedTicket).Return(nil)

		err := service.PurchaseTicket(ticketID, 10)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Insufficient Allocation", func(t *testing.T) {
		ticketID := int16(2)
		initialTicket := models.Ticket{
			Id:         int(ticketID),
			Name:       "example",
			Desc:       "sample description",
			Allocation: 5,
		}

		mockRepo.On("Get", ticketID).Return(initialTicket, nil)

		err := service.PurchaseTicket(ticketID, 10)

		assert.Error(t, err)
		assert.Equal(t, "insufficient allocation", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Ticket Not Found", func(t *testing.T) {
		ticketID := int16(3)

		mockRepo.On("Get", ticketID).Return(models.Ticket{}, errors.New("ticket not found"))

		err := service.PurchaseTicket(ticketID, 10)

		assert.Error(t, err)
		assert.Equal(t, "ticket not found", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update Failure", func(t *testing.T) {
		ticketID := int16(4)
		initialTicket := models.Ticket{
			Id:         int(ticketID),
			Name:       "example",
			Desc:       "sample description",
			Allocation: 100,
		}

		mockRepo.On("Get", ticketID).Return(initialTicket, nil)
		mockRepo.On("Update", mock.Anything).Return(errors.New("failed to update ticket"))

		err := service.PurchaseTicket(ticketID, 10)

		assert.Error(t, err)
		assert.Equal(t, "failed to update ticket", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
