package utils

import (
	"net/http"

	grq "github.com/parnurzeal/gorequest"
)

// RequestHandler
type RequestHandler struct {
}

// IRequestHandler
type IRequestHandler interface {
	ServeRequest(grqAgent *grq.SuperAgent) (response *http.Response, body string, errors []error)
}

// ServeRequest
func (reqestHandler *RequestHandler) ServeRequest(grqAgent *grq.SuperAgent) (response *http.Response, body string, errors []error) {
	return grqAgent.End()
}
