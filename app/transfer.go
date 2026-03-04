package app

import (
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
	}

	helpers.Log(transfer[0].Id, "./logs/transfer.txt")
}
