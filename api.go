package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetGreetingFromAPI sends a request to the API and returns a greeting message
func GetGreetingFromAPI(name string) (string, error) {
	url := fmt.Sprintf("http://localhost:8080/greet?name=%s", name)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get greeting: %d", resp.StatusCode)
	}

	var greeting GreetingResponse
	err = json.NewDecoder(resp.Body).Decode(&greeting)
	if err != nil {
		return "", fmt.Errorf("error parsing response: %v", err)
	}

	return greeting.Message, nil
}
