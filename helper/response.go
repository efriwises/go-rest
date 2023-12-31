package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseWithData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "success"

	if code >= 400 {
		status = "failed"
	}

	if payload != nil {
		response = &ResponseWithData{
			Code:    code,
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Code:    code,
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
