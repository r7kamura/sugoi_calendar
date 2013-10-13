package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseBody struct {
	Message string `json:"message"`
}

type Controller struct {
	Request *http.Request
	ResponseWriter http.ResponseWriter
}

func (controller Controller) RenderJson(statusCode int, data interface{}) {
	controller.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	json, err := json.Marshal(data)
	if err == nil {
		controller.ResponseWriter.WriteHeader(statusCode)
		controller.ResponseWriter.Write(json)
	} else {
		controller.RenderErrorJson(500, "Failed to generage JSON")
	}
}

func (controller Controller) RenderErrorJson(statusCode int, message string) {
	controller.RenderJson(statusCode, ErrorResponseBody{Message: message})
}
