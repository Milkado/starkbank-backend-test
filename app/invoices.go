package app

import (
	"fmt"
	"strconv"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/starkbank/sdk-go/starkbank/invoice"
)

func CreateInvoice() {
	inv := GenerateInvoices()

	invoices, err := invoice.Create(inv, helpers.Auth())

	if err.Errors != nil {
		for _, e := range err.Errors {
			helpers.Log(e.Message, "./logs/stark_error.txt")
		}
		fmt.Println("Error creating invoices")
		return
	}

	fmt.Println("Succerfully created " + strconv.Itoa(len(invoices)) + "invoices")

}
