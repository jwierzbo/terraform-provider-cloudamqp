package console

import (
	"net/http"

	"github.com/dghubble/sling"
)

// API doc: https://docs.cloudamqp.com/cloudamqp_api.html
const customerApiURP = "https://api.cloudamqp.com"

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
