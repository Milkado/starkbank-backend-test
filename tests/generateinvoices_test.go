package tests

import (
	"testing"

	"github.com/Milkado/stark-backend-test/app"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceCount(t *testing.T) {
	invoices := app.GenerateInvoices()

	assert.Equal(t, 10, len(invoices))
}

func TestInvoiceValueRange(t *testing.T) {
	invoices := app.GenerateInvoices()

	clients := app.Clients
	dict := make(map[string]string)

	for _, client := range clients {
		dict[client.Name] = client.Cpf
	}

	for _, inv := range invoices {
		amount := (inv.Amount < 900001 && inv.Amount >= 100000)
		assert.Equal(t, true, amount)
		assert.Equal(t, dict[inv.Name], inv.TaxId)
	}
}
