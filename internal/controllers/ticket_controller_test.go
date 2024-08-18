package controllers_test

import (
	"encoding/json"
	"github.com/akasaa101/ticketing/internal/controllers"
	"github.com/akasaa101/ticketing/internal/models"
	"github.com/akasaa101/ticketing/internal/services/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTicketService(ctrl)
	controller := controllers.NewTicketController(mockService)

	app := fiber.New()
	app.Post("/ticket", controller.CreateTicket)

	t.Run("successful creation with ID as integer", func(t *testing.T) {
		mockTicket := models.Ticket{
			Id:         1,
			Name:       "example",
			Desc:       "sample description",
			Allocation: 100,
		}

		mockService.EXPECT().TicketInsert(gomock.Any()).Return(mockTicket, nil)

		body := `{"name": "example", "desc": "sample description", "allocation": 100}`
		req := httptest.NewRequest(http.MethodPost, "/ticket", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var responseMap map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)

		id, ok := responseMap["id"].(float64)
		assert.True(t, ok, "Expected 'id' to be a number")

		assert.Equal(t, float64(int(id)), id, "Expected 'id' to be an integer")

		assert.Equal(t, "example", responseMap["name"])
		assert.Equal(t, "sample description", responseMap["desc"])
		assert.Equal(t, float64(100), responseMap["allocation"])
	})
	t.Run("missing fields", func(t *testing.T) {
		// No mock expectation here since the service should not be called due to validation failure

		body := `{"name": "", "desc": "", "allocation": -5}` // Invalid data (empty name/desc, negative allocation)
		req := httptest.NewRequest(http.MethodPost, "/ticket", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		// Validate the response status code
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		// Validate the response body contains the error message
		var responseMap map[string]interface{}
		err := json.NewDecoder(resp.Body).Decode(&responseMap)
		assert.NoError(t, err)

		// Assert error response
		assert.Contains(t, responseMap["error"], "Invalid Json")
	})
}
