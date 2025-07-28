package internal

import "fmt"

type HandleFunc func(string) (bool, string)

type Chain struct {
	handlers []HandleFunc
}

func NewChain() *Chain {
	return &Chain{
		handlers: make([]HandleFunc, 0),
	}
}

func (c *Chain) Add(handler HandleFunc) {
	c.handlers = append(c.handlers, handler)
}

func (c *Chain) Process(request string) string {
	for _, handler := range c.handlers {
		if handled, result := handler(request); handled {
			return result
		}
	}

	return "No Handler Found"
}

func CreateAuthHandler() HandleFunc {
	return func(request string) (bool, string) {
		if request == "login" || request == "authenticate" {
			return true, "Authentication successful"
		}
		return false, ""
	}
}

func CreateValidateHandler() HandleFunc {
	return func(request string) (bool, string) {
		if request == "validate" {
			return true, "Validation Successful"
		}
		return false, ""
	}
}

func CreateProcessingHandler() HandleFunc {
	return func(request string) (bool, string) {
		if request == "process" {
			return false, "Processing Success"
		}
		return false, ""
	}
}

func TestFunctionalHandlers() {
	chain := NewChain()

	chain.Add(CreateAuthHandler())
	chain.Add(CreateValidateHandler())
	chain.Add(CreateProcessingHandler())

	request := []string{"login", "validate", "process", "unknown"}

	for _, req := range request {
		res := chain.Process(req)
		fmt.Printf("Request '%s': %s\n", req, res)
	}
}
