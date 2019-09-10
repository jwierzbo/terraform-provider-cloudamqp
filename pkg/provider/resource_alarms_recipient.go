package provider

import (
	"log"

	"github.com/jwierzbo/terraform-provider-cloudamqp/pkg/api/console"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlarmRecipient() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmRecipientCreate,
		Read:   resourceAlarmRecipientRead,
		// This Read method is used only for persisting `console_api_key` which not comes from REST API
		Update: resourceAlarmRecipientRead,
		Delete: resourceAlarmRecipientDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of notifications, valid options are: email, webhook, pagerduty, opsgenie, opsgenie-eu, victorops",
				ForceNew:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Recipient value depended by type",
				ForceNew:    true,
			},
			"console_api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance ApiKey to authorize",
				Sensitive:   true,
			},
		},
	}
}

func resourceAlarmRecipientRead(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	data, err := client.RecipientsList()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Looking for receipient Id: %s", d.Id())

	for _, recipient := range data {
		log.Printf("[DEBUG] Checking receipient: %v", recipient)
		if recipient["id"].(string) == d.Id() {
			for k, v := range recipient {
				_ = d.Set(k, v)
			}
			return nil
		}
	}

	log.Printf("[WARN] Recipient not found: %v", data)
	d.SetId("")
	return nil
}

func resourceAlarmRecipientCreate(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	keys := []string{"type", "value"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}
	data, err := client.RecipientsAdd(params)
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))
	for k, v := range data {
		_ = d.Set(k, v)
	}
	return nil
}

func resourceAlarmRecipientDelete(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	params := map[string]interface{}{
		"id": d.Id(),
	}

	return client.RecipientsDelete(params)
}
