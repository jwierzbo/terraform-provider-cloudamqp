package provider

import (
	"log"

	"github.com/jwierzbo/terraform-provider-cloudamqp/pkg/api/console"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlarmConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmConfigurationCreate,
		Read:   resourceAlarmConfigurationRead,
		// This Read method is used only for persisting `console_api_key` which not comes from REST API
		Update: resourceAlarmRecipientRead,
		Delete: resourceAlarmConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of notifications, valid options are: cpu, memory, disk, queue, connection, consumer, netsplit",
				ForceNew:    true,
			},
			"value_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "What value to trigger the alarm for",
				ForceNew:    true,
			},
			"time_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "For how long (in seconds) the value_threshold should be active before trigger alarm",
				ForceNew:    true,
			},
			"notification_ids": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Optional:    true,
				Description: "Add recipient id to send notification when the alarm is triggered. Leave empty to automatically add all recipients.",
				ForceNew:    true,
			},
			"vhost_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Regex for which vhost the queues are in",
				ForceNew:    true,
			},
			"queue_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Regex for which queues to check",
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

func resourceAlarmConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	data, err := client.AlarmsList()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Looking for alarm Id: %s", d.Id())

	for _, alarm := range data {
		log.Printf("[DEBUG] Checking alarm: %v", alarm)
		if alarm["id"].(string) == d.Id() {
			for k, v := range alarm {
				_ = d.Set(k, v)
			}
			return nil
		}
	}

	log.Printf("[WARN] Alarm not found: %v", data)
	d.SetId("")
	return nil
}

func resourceAlarmConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	keys := []string{"type", "value_threshold", "time_threshold", "notification_ids", "vhost_regex", "queue_regex"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}

	var notificationIDs []int
	if attr := d.Get("notification_ids").(*schema.Set); attr.Len() > 0 {
		for _, v := range attr.List() {
			val, ok := v.(int)
			if ok {
				notificationIDs = append(notificationIDs, val)
			}
		}
	}

	params["notifications"] = notificationIDs
	delete(params, "notification_ids")

	data, err := client.AlarmAdd(params)
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))

	return nil
}

func resourceAlarmConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	client := console.New(d.Get("console_api_key").(string))
	params := map[string]interface{}{
		"alarm_id": d.Id(),
	}

	return client.AlarmDelete(params)
}
