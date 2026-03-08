package app

import (
	"io"
	"net/http"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/labstack/echo/v5"
	Event "github.com/starkbank/sdk-go/starkbank/event"
)

func Listener(c *echo.Context) error {
	bytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		helpers.Log("error reading webhook body", "./logs/error.txt")
		return c.String(http.StatusInternalServerError, "failed to read body")
	}

	resp := WebhookResponseParser(bytes)

	signature := validateSignature(bytes, c.Request().Header.Get("Digital-Signature"))

	if !signature {
		helpers.Log("Signature failed to validate, check error logs", "./logs/error.txt")
		return c.String(http.StatusInternalServerError, "signature validation failed")
	}

	if resp.Event.Log.Type == "credited" {
		TranrferCredited(resp.Event.Log.Invoice.Amount - resp.Event.Log.Invoice.Fee)
	}

	return c.String(http.StatusOK, "listened")
}

func validateSignature(bytes []byte, signature string) bool {
	if signature == "" {
		helpers.Log("Error parsing: empty signature", "./logs/error.txt")
		return false
	}
	_, strkError := Event.Parse(string(bytes), signature, helpers.Auth())

	if strkError.Errors != nil {
		for _, e := range strkError.Errors {
			helpers.Log("Error parsing: "+e.Message, "./logs/error.txt")
		}
		return false
	}

	return true
}
