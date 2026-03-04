package app

import (
	"math/rand/v2"

	"github.com/starkbank/sdk-go/starkbank/invoice"
)

type Client struct {
	Name string
	Cpf  string
}

// Populate the clients slice with names and CPFs
var clients = []Client{
	{Name: "Kaladin Stormblessed", Cpf: "020.663.580-02"},
	{Name: "Cephandrius", Cpf: "141.255.290-78"},
	{Name: "Xisis", Cpf: "729.804.760-48"},
	{Name: "Frost", Cpf: "915.031.780-64"},
	{Name: "Sixth of the Dusk", Cpf: "643.944.520-07"},
	{Name: "Kelsier", Cpf: "074.346.500-87"},
	{Name: "Shallan Davar", Cpf: "362.903.250-86"},
	{Name: "Adolin Kholin", Cpf: "596.316.010-30"},
	{Name: "Jasnah Kholin", Cpf: "256.879.810-63"},
	{Name: "Rlain", Cpf: "875.647.480-60"},
}

func GenerateInvoices() []invoice.Invoice {
	invoices := []invoice.Invoice{}

	for i := 0; i <= 9; i++ {
		// Pick a random client from the slice
		client := clients[rand.IntN(len(clients))]

		amount := rand.IntN(800001) + 100000

		invoices = append(invoices, invoice.Invoice{
			Amount: amount,
			Name:   client.Name,
			TaxId:  client.Cpf,
		})
	}

	return invoices
}
