package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Milkado/stark-backend-test/app"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
)

// TestWebhookResponseParser verifies the JSON unmarshaling logic
func TestWebhookResponseParser(t *testing.T) {
	jsonPayload := []byte(`{
		"event": {
			"log": {
				"type": "credited",
				"invoice": {
					"amount": 50000,
					"status": "paid"
				}
			}
		}
	}`)

	resp := app.WebhookResponseParser(jsonPayload)

	assert.Equal(t, "credited", resp.Event.Log.Type)
	assert.Equal(t, 50000, resp.Event.Log.Invoice.Amount)
}

// TestListener_InvalidSignature tests how the Listener handles a request with a bad signature
func TestListener_InvalidSignature(t *testing.T) {
	e := echo.New()

	body := []byte(`{"test": "data"}`)
	req := httptest.NewRequest(http.MethodPost, "/webhook/payment", bytes.NewBuffer(body))

	dummySignature := "MEYCIQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIhAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	req.Header.Set("Digital-Signature", dummySignature)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := app.Listener(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

// TestListener_EmptyBody verifies error handling for empty requests
func TestListener_EmptyBody(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webhook/payment", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = app.Listener(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
