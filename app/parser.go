package app

import "encoding/json"

type (
	invoiceResp struct {
		Event event `json:"event"`
	}

	event struct {
		Log log `json:"log"`
	}

	log struct {
		Invoice invoiceStatus `json:"invoice"`
		Type    string        `json:"type"`
	}

	invoiceStatus struct {
		Amount        int    `json:"amount"`
		NominalAmount int    `json:"nominalAmount"`
		Fee           int    `json:"fee"`
		Status        string `json:"status"`
	}
)

func WebhookResponseParser(body []byte) invoiceResp {
	resp := new(invoiceResp)

	json.Unmarshal(body, resp)

	return *resp
}
