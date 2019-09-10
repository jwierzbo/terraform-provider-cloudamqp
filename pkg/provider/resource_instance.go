package provider

import (
	"github.com/jwierzbo/terraform-provider-cloudamqp/pkg/api/customer"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		Read:   resourceInstanceRead,
		Update: resourceInstanceUpdate,
		Delete: resourceInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the instance",
			},
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the plan, valid options are: lemur, tiger, bunny, rabbit, panda, ape, hippo, lion",
			},
			"nodes": {
				Type:        schema.TypeInt,
				Default:     1,
				Optional:    true,
				Description: "Number of nodes in cluster (plan must support it)",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the region you want to create your instance in",
			},
			"vpc_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Dedicated VPC subnet, shouldn't overlap with your current VPC's subnet",
			},
			/* 'tags' is not working currently through the REST API (only read operation work)
			"tags": {
				Type:        schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Description: "Comma separated list of tags for instance.",
				ForceNew: false,
			},*/
			"rmq_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "RabbitMQ version",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "URL of the CloudAMQP instance",
			},
			"apikey": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "API key for the CloudAMQP instance",
			},
		},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*customer.API)
	keys := []string{"name", "plan", "region", "nodes"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}

	data, err := client.Create(params)
	if err != nil {
		return err
	}
	d.SetId(data["id"].(string))
	for k, v := range data {
		_ = d.Set(k, v)
	}
	return nil
}

func resourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*customer.API)
	data, err := client.Read(d.Id())
	if err != nil {
		return err
	}
	for k, v := range data {
		_ = d.Set(k, v)
	}
	return nil
}

func resourceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*customer.API)
	keys := []string{"name", "plan", "nodes"}
	params := make(map[string]interface{})
	for _, k := range keys {
		params[k] = d.Get(k)
	}

	return client.Update(d.Id(), params)
}

func resourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*customer.API)
	return client.Delete(d.Id())
}
