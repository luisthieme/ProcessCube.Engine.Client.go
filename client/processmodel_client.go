package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ProcessModelClient struct {
	engineURL string
	basePath string
	identity *Identity
}

func NewProcessModelClient(engineURL string, basePath string, identity *Identity) *ProcessModelClient {
	return &ProcessModelClient{engineURL: engineURL, basePath: basePath, identity: identity}
}


func (c *ProcessModelClient) GetAll() (*ProcessModelsResponse, error) {
	url := c.engineURL + c.basePath + "/process_models"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+ c.identity.Token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", res.StatusCode)
		return nil, err
	}

	var response ProcessModelsResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Printf("failed to decode response: %v", err)
		return nil, err
	}

	return &response, nil

}

func (c *ProcessModelClient) GetById(processModelId string) (*ProcessModel, error) {
	url := c.engineURL + c.basePath + "/process_models/" + processModelId + "/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+ c.identity.Token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		log.Printf("Process definition for ProcessModel `%s` not found!", processModelId)
		err := errors.New("NotFoundError")
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", res.StatusCode)
		return nil, err
	}

	var response ProcessModel
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Printf("failed to decode response: %v", err)
		return nil, err
	}

	return &response, nil
}


