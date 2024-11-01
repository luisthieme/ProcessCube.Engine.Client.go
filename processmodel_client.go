package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type ProcessModelClient struct {
	engineURL string
	basePath string
	identity Identity
}

func NewProcessModels(engineURL string, basePath string, identity Identity) *ProcessModelClient {
	return &ProcessModelClient{engineURL: engineURL, basePath: basePath, identity: identity}
}


func (c *ProcessModelClient) GetAll() {
		url := c.engineURL + c.basePath + "/process_models"
	
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Failed to create request: %v", err)
		}
	
		req.Header.Set("Authorization", "Bearer "+ c.identity.Token)
	
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("Failed to send request: %v", err)
		}
		defer res.Body.Close()
	
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}

	fmt.Println(string(body))
}


