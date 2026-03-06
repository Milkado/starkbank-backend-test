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

// TestListener_InvalidSignature tests how the Listener handles a request with a bad signature
func TestListener_InvalidSignature(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test that interacts with external SDK in short mode")
	}
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
	if testing.Short() {
		t.Skip("skipping test that interacts with external SDK in short mode")
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/webhook/payment", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	_ = app.Listener(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
