package console

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

/* Alarms Recipients */

func (api *API) RecipientsList() ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Get("/api/alarms/recipients").Receive(&data, &failed)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	for index, recipient := range data {
		data[index]["id"] = strconv.FormatFloat(recipient["id"].(float64), 'f', 0, 64 )
	}

	log.Printf("[DEBUG] Returning list of recipients: %v", data)
	return data, nil
}

func (api *API) RecipientsAdd(params map[string]interface{}) (map[string]interface{}, error) {
	log.Printf("[INFO] Creating resource with params: %v", params)
	data := make(map[string]interface{})
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Post("/api/alarms/recipients").BodyJSON(params).Receive(&data, &failed)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	data["id"] = strconv.FormatFloat(data["id"].(float64), 'f', 0, 64 )
	log.Printf("[INFO] Resource created, response: %v", data)
	return data, err
}

func (api *API) RecipientsDelete(params map[string]interface{}) error {
	log.Printf("[INFO] Removing recipient: %v", params)
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Delete("/api/alarms/recipients").BodyJSON(params).Receive(nil, &failed)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}
	return nil
}

/* Alarms Configurations */

func (api *API) AlarmsList() ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Get("/api/alarms").Receive(&data, &failed)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	for index, recipient := range data {
		data[index]["id"] = strconv.FormatFloat(recipient["id"].(float64), 'f', 0, 64 )
	}

	log.Printf("[DEBUG] Returning list of alarms: %v", data)
	return data, nil
}

func (api *API) AlarmAdd(params map[string]interface{}) ( map[string]interface{}, error) {
	log.Printf("[INFO] Creating alarm with params: %v", params)
	data := make(map[string]interface{})
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Post("/api/alarms").BodyJSON(params).Receive(&data, &failed)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	data["id"] = strconv.FormatFloat(data["id"].(float64), 'f', 0, 64 )
	log.Printf("[INFO] Alarm created, response: %v", data)
	return data, err
}

func (api *API) AlarmDelete(params map[string]interface{}) error {
	log.Printf("[INFO] Removing alarm: %v", params)
	failed := make(map[string]interface{})

	resp, err := api.sling.New().Delete("/api/alarms").BodyJSON(params).Receive(nil, &failed)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}
	return nil
}
