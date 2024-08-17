package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/akasaa101/ticketing/internal/database"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/akasaa101/ticketing/internal/handler"
	"github.com/akasaa101/ticketing/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	app := fiber.New()

	app.Post("/tickets", handler.CreateTicket)
	app.Get("/tickets/:id", handler.GetTicket)
	app.Post("/tickets/:id/purchases", handler.PurchaseTicket)

	t.Run("Test CreateTicket", func(t *testing.T) {
		// Create a ticket first
		requestBody := `{"name":"example","desc":"sample description","allocation":100}`
		req := httptest.NewRequest("POST", "/tickets", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)

		id, ok := responseBody["id"].(float64)
		assert.True(t, ok)
		assert.True(t, id > 0)

		expectedResponse := map[string]interface{}{
			"name":       "example",
			"desc":       "sample description",
			"allocation": float64(100),
		}

		delete(responseBody, "id")

		assert.Equal(t, expectedResponse, responseBody)

		// Check if the record exists in the database
		db := database.DB.Db
		var ticket model.Ticket

		db.Find(&ticket, "id = ?", int(id))

		assert.Equal(t, int(id), ticket.ID)
		assert.Equal(t, "example", ticket.Name)
		assert.Equal(t, "sample description", ticket.Desc)
		assert.Equal(t, 100, ticket.Allocation)
	})

	t.Run("Test GetTicket", func(t *testing.T) {
		// Create a ticket first
		createRequestBody := `{"name":"example","desc":"sample description","allocation":100}`
		createReq := httptest.NewRequest("POST", "/tickets", strings.NewReader(createRequestBody))
		createReq.Header.Set("Content-Type", "application/json")
		createResp, err := app.Test(createReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, createResp.StatusCode)

		var createResponseBody map[string]interface{}
		err = json.NewDecoder(createResp.Body).Decode(&createResponseBody)
		assert.NoError(t, err)
		ticketID, ok := createResponseBody["id"].(float64)
		assert.True(t, ok)
		assert.True(t, ticketID > 0)

		// Retrieve the ticket via GET request
		getReq := httptest.NewRequest("GET", "/tickets/"+strconv.Itoa(int(ticketID)), nil)
		getResp, err := app.Test(getReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, getResp.StatusCode)

		var getResponseBody map[string]interface{}
		err = json.NewDecoder(getResp.Body).Decode(&getResponseBody)
		assert.NoError(t, err)

		expectedResponse := map[string]interface{}{
			"id":         ticketID,
			"name":       "example",
			"desc":       "sample description",
			"allocation": float64(100),
		}

		assert.Equal(t, expectedResponse, getResponseBody)

		// Verify the ticket exists in the database and matches the expected values
		db := database.DB.Db
		var ticket model.Ticket
		err = db.Find(&ticket, "id = ?", int(ticketID)).Error
		assert.NoError(t, err)

		// Assert that the ticket in the database matches what was created
		assert.Equal(t, "example", ticket.Name)
		assert.Equal(t, "sample description", ticket.Desc)
		assert.Equal(t, 100, ticket.Allocation)
		assert.Equal(t, int(ticketID), ticket.ID)
	})

	t.Run("Test PurchaseTicket", func(t *testing.T) {
		// Create a ticket first
		requestBody := `{"name":"example","desc":"sample description","allocation":100}`
		req := httptest.NewRequest("POST", "/tickets", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&createResponse)
		assert.NoError(t, err)

		ticketID, ok := createResponse["id"].(float64)
		assert.True(t, ok)

		// Purchase a ticket
		purchaseRequestBody := `{"quantity": 10}`
		purchaseReq := httptest.NewRequest("POST", fmt.Sprintf("/tickets/%d/purchase", int(ticketID)), strings.NewReader(purchaseRequestBody))
		purchaseReq.Header.Set("Content-Type", "application/json")
		purchaseResp, err := app.Test(purchaseReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, purchaseResp.StatusCode)

		// Check if the ticket allocation has been updated in the database
		db := database.DB.Db
		var ticket model.Ticket
		db.Find(&ticket, "id = ?", int(ticketID))

		assert.Equal(t, 90, ticket.Allocation)
	})

	t.Run("Test PurchaseTicketInsufficientAllocation", func(t *testing.T) {
		// Create a ticket with low allocation
		requestBody := `{"name":"example","desc":"sample description","allocation":5}`
		req := httptest.NewRequest("POST", "/tickets", strings.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&createResponse)
		assert.NoError(t, err)

		ticketID, ok := createResponse["id"].(float64)
		assert.True(t, ok)

		// Try to purchase more tickets than available
		purchaseRequestBody := `{"quantity": 10}`
		purchaseReq := httptest.NewRequest("POST", fmt.Sprintf("/tickets/%d/purchase", int(ticketID)), strings.NewReader(purchaseRequestBody))
		purchaseReq.Header.Set("Content-Type", "application/json")
		purchaseResp, err := app.Test(purchaseReq)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, purchaseResp.StatusCode)

		// Check if the ticket allocation has remained the same in the database
		db := database.DB.Db
		var ticket model.Ticket
		db.Find(&ticket, "id = ?", int(ticketID))

		assert.Equal(t, 5, ticket.Allocation)
	})

}
