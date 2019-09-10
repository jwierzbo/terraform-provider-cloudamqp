package customer

import (
	"net/http"

	"github.com/dghubble/sling"
)

// API doc: https://docs.cloudamqp.com/
const customerApiURP = "https://customer.cloudamqp.com"

type API struct {
	sling *sling.Sling
}

func New(apiKey string) *API {
	return &API{
		sling: sling.New().
			Client(http.DefaultClient).
			Base(customerApiURP).
			SetBasicAuth("", apiKey),
	}
}
