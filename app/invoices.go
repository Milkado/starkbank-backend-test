package app

import (
	"fmt"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/starkbank/sdk-go/starkbank/invoice"
)

func CreateInvoice() {
	inv := GenerateInvoices()

	invoices, err := invoice.Create(inv, helpers.Auth())

	if err.Errors != nil {
		for _, e := range err.Errors {
			helpers.Log(e.Message, "./logs/error.txt")
		}
		fmt.Println("Error creating invoices")
		return
	}

	for _, inv := range invoices {
		amountFloat := float64(inv.Amount) / 100.0
		helpers.Log("Invoice "+inv.Id+" created for "+inv.Name+" with amount "+fmt.Sprintf("%.2f", amountFloat), "./logs/created.txt")
	}
}
