package main

import (
	"github.com/jwierzbo/terraform-provider-cloudamqp/pkg/provider"

	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
