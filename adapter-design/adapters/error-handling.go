package adapters

import (
	"errors"
	"fmt"
)

type ApiClient interface {
	MakeRequest(endPoint string, data map[string]interface{}) (map[string]interface{}, error)
	GetName() string
}

// Adaptee - Legacy Adaptee
type ThirdPartyService struct{}

func (t *ThirdPartyService) sendRequest(url string, payload map[string]interface{}) (interface{}, error) {
	fmt.Println("The url is : ", url)
	if url == "/timeout" {
		return nil, errors.New("request timeout")
	}

	if url == "/error" {
		return nil, errors.New("service error")
	}

	fmt.Printf("Sending request to %s with payload: %+v\n", url, payload)

	return map[string]interface{}{
		"status": "success",
		"data":   "Response from third party",
	}, nil
}

func (t *ThirdPartyService) GetServiceName() string {
	return "ThirdPartyAPI"
}

// Adapter

type ThirdPartyAdapter struct {
	service *ThirdPartyService
}

func NewThirdPartyAdapter(srv *ThirdPartyService) *ThirdPartyAdapter {
	return &ThirdPartyAdapter{
		service: srv,
	}
}

// Implement client methods
// Here delegate the request to other service

func (t ThirdPartyAdapter) MakeRequest(endPoint string, data map[string]interface{}) (map[string]interface{}, error) {
	// make an api call to service sendRequest method and get Response

	// url := fmt.Sprintf("/api/v1%s", endPoint)

	response, err := t.service.sendRequest(endPoint, data)

	if err != nil {
		switch err.Error() {
		case "request timeout":
			return nil, fmt.Errorf("adapter_timeout: request to %s timed out ", t.GetName())
		case "service error":
			return nil, fmt.Errorf("server error : %s ", t.GetName())
		default:
			return nil, fmt.Errorf("adapter_error: %w", err)
		}
	}

	if respMap, ok := response.(map[string]interface{}); ok {
		return respMap, nil
	}

	return nil, errors.New("invalid Response format")
}

func (t ThirdPartyAdapter) GetName() string {
	return t.service.GetServiceName()
}

// Client with err handling

func ProcessRequestHandling(client ApiClient, endPoint string, data map[string]interface{}) {
	fmt.Printf("Making request to %s...\n", client.GetName())

	response, err := client.MakeRequest(endPoint, data)

	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}
	fmt.Printf("Success: %+v\n", response)
}

func TestAdapterErrorHandling() {
	service := &ThirdPartyService{}
	adp := NewThirdPartyAdapter(service)

	// Request call

	ProcessRequestHandling(adp, "/users", map[string]interface{}{
		"name": "John",
		"age":  30,
	})

	// Error cases
	ProcessRequestHandling(adp, "/error", nil)
	ProcessRequestHandling(adp, "/timeout", nil)
}
