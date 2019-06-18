package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/miliyash/ms-contact-manager/app/utils"
	log "github.com/miliyash/ms-contact-manager/logging"
)

const jsonContentType = "application/json"

// ERRORS ...
const ERRORS = "errors"

// JSON => Add content type to Json and Marshel Json header data
func JSON(c context.Context, w http.ResponseWriter, status int, v interface{}) {
	w.Header().Add("Content-Type", jsonContentType)

	ba, err := json.Marshal(v)
	if err != nil {
		log.WithError(err).Error("basecontroller.JSON failed to marshal")
	}
	w.WriteHeader(status)
	_, _ = w.Write(ba)

}

// HandleGetResponse marshals the result and renders it to the response
func HandleGetResponse(c context.Context, w http.ResponseWriter, err *utils.JSONAPIError, result utils.Renderable) {
	logError(c, err)
	var j []byte
	var jsonError error
	var response interface{}

	if err == nil {
		if j, jsonError = json.MarshalIndent(result, "", "  "); jsonError != nil {
			JSON(c, w, 400, map[string]interface{}{ERRORS: jsonError})
		}
		if jsonError = json.Unmarshal(j, &response); jsonError != nil {
			JSON(c, w, 400, map[string]interface{}{ERRORS: jsonError})
		}
		JSON(c, w, 200, response)
	} else {
		JSON(c, w, 400, map[string]interface{}{ERRORS: err})
	}
}

// HandlePostResponse => wraps json-api handler
func HandlePostResponse(c context.Context, w http.ResponseWriter, success bool, err *utils.JSONAPIError, resource utils.Renderable) {
	logError(c, err)
	var j []byte
	var jsonError error
	var response interface{}

	if success {
		if j, jsonError = json.MarshalIndent(resource, "", "  "); jsonError != nil {
			JSON(c, w, 400, map[string]interface{}{ERRORS: jsonError})
		}
		if jsonError = json.Unmarshal(j, &response); jsonError != nil {
			JSON(c, w, 400, map[string]interface{}{ERRORS: jsonError})
		}
		JSON(c, w, 201, response)
	} else if err != nil {
		// TODO: how do I parse the status code?
		JSON(c, w, 400, map[string]interface{}{ERRORS: err})
	} else {
		JSON(c, w, 422, map[string]interface{}{ERRORS: resource.Errors()})
	}
}

func logError(c context.Context, err *utils.JSONAPIError) {
	if err != nil {
		// fields := map[string]interface{}{"err": err.Detail}
		// utils.LogError(err.Detail, &fields)
		log.Error(err.Detail, nil)
	}
}
