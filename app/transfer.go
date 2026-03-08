package app

import (
	"fmt"

	"github.com/Milkado/stark-backend-test/helpers"
	"github.com/starkbank/sdk-go/starkbank/transfer"
)

func TranrferCredited(amount int) {
	transfer, err := transfer.Create(
		[]transfer.Transfer{
			{
				Amount:        amount,
				Name:          helpers.Env("NAME"),
				TaxId:         helpers.Env("TAX_ID"),
				BankCode:      helpers.Env("BANK_CODE"),
				BranchCode:    helpers.Env("BRANCH"),
				AccountNumber: helpers.Env("ACCOUNT"),
				AccountType:   helpers.Env("ACCOUNT_TYPE"),
			},
		}, helpers.Auth())

	if err.Errors != nil {
		for _, e := range err.Errors {
			helpers.Log("Transfer error: "+e.Message, "./logs/error.txt")
		}
		return
	}

	t := transfer[0]
	amountFloat := float64(t.Amount) / 100.0
	helpers.Log(t.Id+" to "+t.Name+" for amount "+fmt.Sprintf("%.2f", amountFloat), "./logs/transfer.txt")
}
