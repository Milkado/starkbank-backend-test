package tests

import (
	"testing"

	"github.com/Milkado/stark-backend-test/app"
	"github.com/stretchr/testify/assert"
)

func TestParser_ValidJson(t *testing.T) {
	jsonPayload := []byte(`{
		"event": {
			"log": {
				"type": "credited",
				"invoice": {
					"amount": 50000,
					"nominalAmount": 49500,
					"fee": 500,
					"status": "paid"
				}
			}
		}
	}`)

	resp := app.WebhookResponseParser(jsonPayload)

	assert.Equal(t, "credited", resp.Event.Log.Type)
	assert.Equal(t, 50000, resp.Event.Log.Invoice.Amount)
	assert.Equal(t, 49500, resp.Event.Log.Invoice.NominalAmount)
	assert.Equal(t, 500, resp.Event.Log.Invoice.Fee)
	assert.Equal(t, "paid", resp.Event.Log.Invoice.Status)
}

func TestParser_InvalidJson(t *testing.T) {
	// Malformed JSON (missing closing brace and quotes)
	invalidPayload := []byte(`{"event": {"log": {"type": "credited"`)

	resp := app.WebhookResponseParser(invalidPayload)

	// In Go, json.Unmarshal returns an error which is ignored in WebhookResponseParser.
	// We expect a zeroed-out struct or default values for uninitialized fields.
	assert.Equal(t, "", resp.Event.Log.Type)
	assert.Equal(t, 0, resp.Event.Log.Invoice.Amount)
}
