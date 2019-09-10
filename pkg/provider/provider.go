package provider

import (
	"github.com/jwierzbo/terraform-provider-cloudamqp/pkg/api/customer"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
)

// Provider returns a new CloudAMQP provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"customer_api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDAMQP_CUSTOMER_APIKEY", nil),
				Description: "Key used to create and edit Instances for CloudAMQP",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudamqp_instance":            resourceInstance(),
			"cloudamqp_alarm_recipient":     resourceAlarmRecipient(),
			"cloudamqp_alarm_configuration": resourceAlarmConfiguration(),
		},
		ConfigureFunc: getCustomerClient,
	}
}

func getCustomerClient(rd *schema.ResourceData) (interface{}, error) {
	customerApiKey, ok := rd.Get("customer_api_key").(string)
	if !ok {
		return nil, errors.New("invalid type for customer_api_key")
	}

	if customerApiKey == "" {
		return nil, errors.New("customerApiKey can not be empty")
	}

	return customer.New(customerApiKey), nil
}
