package internal

import "fmt"

// Handler interface
type Handler interface {
	SetNext(handle Handler) Handler
	Handle(request string) string
}

// Base handle

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handle Handler) Handler {
	b.next = handle
	return b.next
}

func (b *BaseHandler) Handle(request string) string {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return ""
}

// Concrete Handlers

type AuthenticationHandle struct {
	BaseHandler
}

func (a *AuthenticationHandle) Handle(request string) string {
	if request == "authenticate" {
		return "User Authenticated successfully"
	}
	return a.next.Handle(request)
}

type AuthorizationHandle struct {
	BaseHandler
}

func (a *AuthorizationHandle) Handle(request string) string {
	if request == "authorize" {
		return "User Authorization Successfully"
	}
	return a.next.Handle(request)
}

type ValidationHandler struct {
	BaseHandler
}

func (h *ValidationHandler) Handle(request string) string {
	if request == "validate" {
		return "Data validated successfully"
	}
	return h.BaseHandler.Handle(request)
}

type ProcessingHandler struct {
	BaseHandler
}

func (h *ProcessingHandler) Handle(request string) string {
	if request == "process" {
		return "Request processed successfully"
	}
	return h.BaseHandler.Handle(request)
}

func TestHandler() {
	authHandler := &AuthenticationHandle{}
	authorizationHandler := &AuthorizationHandle{}
	validationHandler := &ValidationHandler{}
	processingHandler := &ProcessingHandler{}

	authHandler.SetNext(authorizationHandler).SetNext(validationHandler).SetNext(processingHandler)

	req := []string{"authenticate", "authorize", "validate", "process", "unknown"}

	for _, v := range req {
		res := authHandler.Handle(v)

		if res == "" {
			fmt.Printf("Request '%s': No handler found\n", req)
		} else {
			fmt.Printf("Request '%s': %s\n", req, res)
		}
	}

}

// logger is a bit extended version of this
