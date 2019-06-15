package utils

import (
	"net/http"

	grq "github.com/parnurzeal/gorequest"
)

// RequestHandler => Handle nlog request and serve response
type RequestHandler struct {
}

// IRequestHandler => Handle nlog request and serve response
type IRequestHandler interface {
	ServeRequest(grqAgent *grq.SuperAgent) (response *http.Response, body string, errors []error)
}

// ServeRequest => Handle nlog Http Serve Request
func (reqestHandler *RequestHandler) ServeRequest(grqAgent *grq.SuperAgent) (response *http.Response, body string, errors []error) {
	return grqAgent.End()
}
