package customer

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (api *API) waitUntilReady(id string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})

	for {
		resp, err := api.sling.Path("/api/instances/").Get(id).Receive(&data, &failed)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
		}

		if data["ready"] == true {
			data["id"] = id
			return data, nil
		}
		time.Sleep(10 * time.Second)
	}
}

func (api *API) Create(params map[string]interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	failed := make(map[string]interface{})

	resp, err := api.sling.Post("/api/instances").BodyJSON(params).Receive(&data, &failed)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	stringId := strconv.Itoa(int(data["id"].(float64)))
	return api.waitUntilReady(stringId)
}

func (api *API) Read(id string) (map[string]interface{}, error) {
	log.Printf("[DEBUG] Getting instance: %s", id)
	data := make(map[string]interface{})
	failed := make(map[string]interface{})

	resp, err := api.sling.Path("/api/instances/").Get(id).Receive(&data, &failed)
	if err != nil {
		log.Printf("[ERROR] Get instance error - Message: %s", failed)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	log.Printf("[DEBUG] Returning instance: %v", data)

	return data, nil
}

func (api *API) Update(id string, params map[string]interface{}) error {
	failed := make(map[string]interface{})

	resp, err := api.sling.Put("/api/instances/" + id).BodyJSON(params).Receive(nil, &failed)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	return err
}

func (api *API) Delete(id string) error {
	failed := make(map[string]interface{})

	resp, err := api.sling.Path("/api/instances/").Delete(id).Receive(nil, &failed)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return errors.New(fmt.Sprintf("Wrong response Status: %v, Message: %s", resp.Status, failed))
	}

	return err
}
